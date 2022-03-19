package repos

import (
	"context"
	"database/sql"
	"fmt"

	"k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

const (
	CustomerLoginError  = "incorrect username and password combination"
	CustomerUpdateError = "can't update customer, missing customer ID"
)

func newCustomer(repo *repo) Customer {
	return &customerRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type Customer interface {
	Exists() bool
	Deleted() bool
	Create(customer *models.Customer) error
	Update(customer *models.Customer) error
	Save() error
	Upsert() error
	Delete(id int64) error
	Login(storeId int64, email, password string) (res *models.Customer, err error)
	GetCustomerByEmail(storeId int64, email string) (res *models.Customer, err error)
}

type customerRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.Customer
}

func (m *customerRepo) Login(storeId int64, email, password string) (res *models.Customer, err error) {
	res, err = models.CustomerByStoreIDEmail(m.ctx, m.db, storeId, email)
	if err != nil {
		return nil, &RepoError{Err: err}
	}

	if m.checkPasswordHash(password, res.Password) {
		return res, nil
	}

	return nil, &RepoError{
		Err:        fmt.Errorf(CustomerLoginError),
		StatusCode: CustomerLoginErrorCode,
	}
}

func (m *customerRepo) Create(customer *models.Customer) error {
	// hash the password
	if customer.Password.Valid {
		hash, _ := m.hashPassword(customer.Password.String)
		customer.Password = sql.NullString{String: hash, Valid: true}
	}
	if err := customer.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *customerRepo) Update(customer *models.Customer) error {
	if customer.ID == 0 {
		return &RepoError{
			Err:        fmt.Errorf(CustomerUpdateError),
			StatusCode: UpdateErrorCode,
		}
	}
	if err := customer.Update(m.ctx, m.db); err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: UpdateErrorCode,
		}
	}
	return nil
}

func (m *customerRepo) Save() error {
	if err := m.Customer.Save(m.ctx, m.db); err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: SaveErrorCode,
		}
	}
	return nil
}

func (m *customerRepo) Upsert() error {
	if err := m.Customer.Upsert(m.ctx, m.db); err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: UpsertErrorCode,
		}
	}
	return nil
}

func (m *customerRepo) Delete(id int64) error {
	customer, err := models.CustomerByID(m.ctx, m.db, id)
	if err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: DeleteErrorCode,
		}
	}
	if err := customer.Delete(m.ctx, m.db); err != nil {
		return &RepoError{
			Err:        err,
			StatusCode: DeleteErrorCode,
		}
	}
	return nil
}

func (m *customerRepo) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		err = &RepoError{
			Err:        err,
			StatusCode: HashPasswordErrorCode,
		}
	}
	return string(bytes), err
}

func (m *customerRepo) checkPasswordHash(password string, hash sql.NullString) bool {
	if hash.Valid {
		err := bcrypt.CompareHashAndPassword([]byte(hash.String), []byte(password))
		return err == nil
	}
	return false
}

func (m *customerRepo) GetCustomerByEmail(storeId int64, email string) (*models.Customer, error) {
	customer, err := models.CustomerByStoreIDEmail(m.ctx, m.db, storeId, email)
	if err != nil {
		err = &RepoError{
			Err:        err,
			StatusCode: GetCustomerByEmailErrorCode,
		}
	}
	return customer, err
}
