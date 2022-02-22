package repos

import (
	"context"
	"fmt"

	"k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func newInventoryStock(repo *repo) InventoryStock {
	return &inventoryStockRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

//go:generate mockgen -destination=./mocks/InventoryStock.go -package=mock_repos k8scommerce/internal/repos InventoryStock
type InventoryStock interface {
	Exists() bool
	Deleted() bool
	Create(ii *models.InventoryStock) error
	Update(ii *models.InventoryStock) error
	Save() error
	Upsert() error
	Delete(id int64) error
}

type inventoryStockRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.InventoryStock
}

func (m *inventoryStockRepo) Create(ii *models.InventoryStock) error {
	// hash the password
	if err := ii.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *inventoryStockRepo) Update(ii *models.InventoryStock) error {
	if ii.ID == 0 {
		return fmt.Errorf("error: can't update ii, missing ii ID")
	}
	if err := ii.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *inventoryStockRepo) Save() error {
	return m.InventoryStock.Save(m.ctx, m.db)
}

func (m *inventoryStockRepo) Upsert() error {
	return m.InventoryStock.Upsert(m.ctx, m.db)
}

func (m *inventoryStockRepo) Delete(id int64) error {
	ii, err := models.ProductByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return ii.Delete(m.ctx, m.db)
}

func (m *inventoryStockRepo) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (m *inventoryStockRepo) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
