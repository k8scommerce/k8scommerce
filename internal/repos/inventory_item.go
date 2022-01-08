package repos

import (
	"context"
	"fmt"

	"k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func newInventoryItem(repo *repo) InventoryItem {
	return &inventoryItemRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type InventoryItem interface {
	Exists() bool
	Deleted() bool
	Create(ii *models.InventoryItem) error
	Update(ii *models.InventoryItem) error
	Save() error
	Upsert() error
	Delete(id int64) error
}

type inventoryItemRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.InventoryItem
}

func (m *inventoryItemRepo) Create(ii *models.InventoryItem) error {
	// hash the password
	if err := ii.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *inventoryItemRepo) Update(ii *models.InventoryItem) error {
	if ii.ID == 0 {
		return fmt.Errorf("error: can't update ii, missing ii ID")
	}
	if err := ii.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *inventoryItemRepo) Save() error {
	return m.InventoryItem.Save(m.ctx, m.db)
}

func (m *inventoryItemRepo) Upsert() error {
	return m.InventoryItem.Upsert(m.ctx, m.db)
}

func (m *inventoryItemRepo) Delete(id int64) error {
	ii, err := models.ProductByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return ii.Delete(m.ctx, m.db)
}

func (m *inventoryItemRepo) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (m *inventoryItemRepo) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
