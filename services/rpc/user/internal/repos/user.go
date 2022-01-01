package repos

import (
	"context"
	"fmt"

	"github.com/k8s-commerce/k8s-commerce/pkg/models"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func newUser(repo *repo) User {
	return &userRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type User interface {
	Create(user *models.User) error
	Update(user *models.User) error
	Save() error
	Upsert() error
	Delete(id int64) error
	Login(username, password string) (res *models.User, err error)
}

type userRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.User
}

func (m *userRepo) Login(username, password string) (res *models.User, err error) {
	res, err = models.UserByEmail(m.ctx, m.db, username)
	if err != nil {
		return nil, err
	}

	if m.checkPasswordHash(password, res.Password) {
		return res, nil
	}

	return nil, fmt.Errorf("error: incorrect username and password combination")
}

func (m *userRepo) Create(user *models.User) error {
	// hash the password
	hash, _ := m.hashPassword(user.Password)
	user.Password = hash

	if err := user.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *userRepo) Update(user *models.User) error {
	if user.ID == 0 {
		return fmt.Errorf("error: can't update user, missing user ID")
	}
	if err := user.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *userRepo) Save() error {
	return m.User.Save(m.ctx, m.db)
}

func (m *userRepo) Upsert() error {
	return m.User.Upsert(m.ctx, m.db)
}

func (m *userRepo) Delete(id int64) error {
	user, err := models.ProductByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return user.Delete(m.ctx, m.db)
}

func (m *userRepo) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (m *userRepo) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
