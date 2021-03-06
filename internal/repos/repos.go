package repos

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/logx"

	_ "github.com/jackc/pgx/v4/stdlib" //postgres driver for sqlx
	"github.com/jmoiron/sqlx"
)

var (
	ErrNotFound = sql.ErrNoRows
)

func NewRepo(c *PostgresConfig) Repo {
	// catch any panics
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("Panic Recovered in NewRepo:", rec)
		}
	}()

	r := &repo{
		cfg: c,
	}

	r.mu.Lock()
	r.db = r.mustConnect()
	r.mu.Unlock()

	return r
}

type Repo interface {
	GetRawDB() *sqlx.DB
	// Repos
	Asset() Asset
	Cart() Cart
	CartItem() CartItem
	Category() Category
	Customer() Customer
	CustomerAddress() CustomerAddress
	CustomerPasswordReset() CustomerPasswordReset
	InventoryBrand() InventoryBrand
	InventoryItem() InventoryItem
	InventoryStock() InventoryStock
	InventorySupplier() InventorySupplier
	Product() Product
	OthersBought() OthersBought
	SimilarProducts() SimilarProducts
	Store() Store
	StoreSetting() StoreSetting
	User() User

	Begin() (*sql.Tx, error)
	Rollback(tx *sql.Tx) error
	Commit(tx *sql.Tx) error
}

type repo struct {
	db       *sqlx.DB
	listener *pq.Listener
	cfg      *PostgresConfig
	mu       sync.Mutex
	once     sync.Once
}

func (r *repo) GetRawDB() *sqlx.DB {
	return r.db
}

func (r *repo) Begin() (*sql.Tx, error) {
	return r.db.Begin()
}

func (r *repo) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

func (r *repo) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

// Repos
func (r *repo) Asset() Asset {
	return newAsset(r)
}

func (r *repo) Cart() Cart {
	return newCart(r)
}

func (r *repo) CartItem() CartItem {
	return newCartItem(r)
}

func (r *repo) Category() Category {
	return newCategory(r)
}

func (r *repo) CustomerAddress() CustomerAddress {
	return newCustomerAddress(r)
}

func (r *repo) Customer() Customer {
	return newCustomer(r)
}

func (r *repo) CustomerPasswordReset() CustomerPasswordReset {
	return newCustomerPasswordReset(r)
}

func (r *repo) InventoryBrand() InventoryBrand {
	return newInventoryBrand(r)
}

func (r *repo) InventoryItem() InventoryItem {
	return newInventoryItem(r)
}

func (r *repo) InventoryStock() InventoryStock {
	return newInventoryStock(r)
}

func (r *repo) InventorySupplier() InventorySupplier {
	return newInventorySupplier(r)
}

func (r *repo) Product() Product {
	return newProduct(r)
}

func (r *repo) OthersBought() OthersBought {
	return newOthersBought(r)
}

func (r *repo) SimilarProducts() SimilarProducts {
	return newSimilarProducts(r)
}

func (r *repo) Store() Store {
	return newStore(r)
}

func (r *repo) StoreSetting() StoreSetting {
	return newStoreSetting(r)
}

func (r *repo) User() User {
	return newUser(r)
}

// connection
func (a *repo) mustConnect() (conn *sqlx.DB) {
	// create the db client
	a.once.Do(func() {
		go a.setDBListener(a.cfg.DataSourceName)
		logx.Info("DB Connecting", "conn", a.cfg.DataSourceName)

		a.dbConnect()

		go func() {
			for {
				a.mu.Lock()
				err := a.db.Ping()
				if err != nil {
					logx.Info("DB Ping failed: ", err)
				}
				a.mu.Unlock()
				time.Sleep(10 * time.Second)
			}
		}()
	})

	return a.db
}

func (a *repo) dbConnect() {
	var conn *sqlx.DB
	conn, err := sqlx.Connect("pgx", a.cfg.DataSourceName)
	if err != nil {
		panic(err)
	}
	a.db = conn
	a.db.SetMaxOpenConns(a.cfg.MaxOpenConnections)
	a.db.SetMaxIdleConns(a.cfg.MaxIdleConnections)
	a.db.SetConnMaxLifetime(time.Minute * time.Duration(a.cfg.MaxConnectionLifetimeMinutes))
	logx.Info("DB status:", "CONNECTED")

	// "kick" the channel to ensure it checks for records on startup
	go a.kickDBChannel()
}

func (a *repo) setDBListener(dburl string) {
	a.listener = pq.NewListener(dburl, 10*time.Second, time.Minute, func(event pq.ListenerEventType, err error) {
		if err != nil {
			logx.Info("DB LISTENER ERROR: ", err)
		}
	})

	if err := a.listener.Listen("raise_notice"); err != nil {
		logx.Info("DB LISTENER ERROR: ", err)
	}

	for {
		select {
		case n := <-a.listener.Notify:
			// n.Extra contains the payload from the notification
			if n != nil {
				// connect here
				// logx.Info(fmt.Sprintf("DB: %s", n.Extra))
			}
		case <-time.After(15 * time.Second):
			if err := a.listener.Ping(); err != nil {
				logx.Info("DB LISTENER ERROR: ", err)
			}
		}
	}
}

func (a *repo) kickDBChannel() {
	a.mu.Lock()
	defer a.mu.Unlock()

	// `pg_notify('commission_payment_jobs_channel'`
	_, err := a.db.Exec(`SELECT pg_notify('commission_payment_jobs_channel','worker startup')`)
	if err != nil {
		logx.Info("DB KICK ERROR: ", err)
	}
}
