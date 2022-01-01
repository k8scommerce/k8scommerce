package repos

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/k8scommerce/k8scommerce/services/rpc/cart/internal/config"

	"github.com/lib/pq"
	"github.com/tal-tech/go-zero/core/logx"

	_ "github.com/jackc/pgx/v4/stdlib" //postgres driver for sqlx
	"github.com/jmoiron/sqlx"
)

var (
	once        sync.Once
	ErrNotFound = sql.ErrNoRows
)

func MustNewRepo(c *config.Config) Repo {
	// catch any panics
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("Panic Recovered in MustNewRepo:", rec)
		}
	}()

	r := &repo{
		cfg: c,
	}
	r.db = r.mustConnect()

	return r
}

type Repo interface {
	GetRawDB() *sqlx.DB
	// Repos
	Cart() Cart
	CartItem() CartItem
	Product() Product
	User() User
}

type repo struct {
	db       *sqlx.DB
	listener *pq.Listener
	cfg      *config.Config
}

func (r *repo) GetRawDB() *sqlx.DB {
	return r.db
}

// Repos
func (r *repo) Cart() Cart {
	return newCart(r)
}

func (r *repo) CartItem() CartItem {
	return newCartItem(r)
}

func (r *repo) Product() Product {
	return newProduct(r)
}

func (r *repo) User() User {
	return newUser(r)
}

// connection
func (a *repo) mustConnect() (conn *sqlx.DB) {
	// create the db client
	once.Do(func() {
		go a.setDBListener(a.cfg.Postgres.Connection)
		logx.Info("DB Connecting", "conn", a.cfg.Postgres.Connection)

		var conn *sqlx.DB
		conn, err := sqlx.Connect("pgx", a.cfg.Postgres.Connection)
		if err != nil {
			panic(err)
		}
		a.db = conn
		a.db.SetMaxOpenConns(a.cfg.Postgres.MaxOpenConnections)
		a.db.SetMaxIdleConns(a.cfg.Postgres.MaxIdleConnections)
		a.db.SetConnMaxLifetime(time.Minute * time.Duration(a.cfg.Postgres.MaxConnectionLifetimeMinutes))
		logx.Info("DB status:", "CONNECTED")

		// "kick" the channel to ensure it checks for records on startup
		a.kickDBChannel()

		go func() {
			for {
				err := a.db.Ping()
				if err != nil {
					logx.Info("DB Ping failed: ", err)
				} else {
					// logx.Info("Pinged successfully")
				}

				time.Sleep(10 * time.Second)
			}
		}()
	})

	return a.db
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
	// `pg_notify('commission_payment_jobs_channel'`
	_, err := a.db.Exec(`SELECT pg_notify('commission_payment_jobs_channel','worker startup')`)
	if err != nil {
		logx.Info("DB KICK ERROR: ", err)
	}
}
