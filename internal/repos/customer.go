package repos

import (
	"context"
	"fmt"

	"k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
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
	Create(user *models.Customer) error
	Update(user *models.Customer) error
	Save() error
	Upsert() error
	Delete(id int64) error
	Login(username, password string) (res *models.Customer, err error)
}

type customerRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.Customer
}

func (m *customerRepo) Login(username, password string) (res *models.Customer, err error) {
	res, err = models.CustomerByEmail(m.ctx, m.db, username)
	if err != nil {
		return nil, err
	}

	if m.checkPasswordHash(password, res.Password) {
		return res, nil
	}

	return nil, fmt.Errorf("error: incorrect username and password combination")
}

func (m *customerRepo) Create(user *models.Customer) error {
	// hash the password
	hash, _ := m.hashPassword(user.Password)
	user.Password = hash

	if err := user.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *customerRepo) Update(user *models.Customer) error {
	if user.ID == 0 {
		return fmt.Errorf("error: can't update user, missing user ID")
	}
	if err := user.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *customerRepo) Save() error {
	return m.Customer.Save(m.ctx, m.db)
}

func (m *customerRepo) Upsert() error {
	return m.Customer.Upsert(m.ctx, m.db)
}

func (m *customerRepo) Delete(id int64) error {
	user, err := models.ProductByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return user.Delete(m.ctx, m.db)
}

func (m *customerRepo) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (m *customerRepo) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
