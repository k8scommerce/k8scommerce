package repos

import (
	"context"
	"fmt"

	"github.com/k8scommerce/k8scommerce/pkg/models"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func newInventorySupplier(repo *repo) InventorySupplier {
	return &inventorySupplierRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type InventorySupplier interface {
	Exists() bool
	Deleted() bool
	Create(ii *models.InventorySupplier) error
	Update(ii *models.InventorySupplier) error
	Save() error
	Upsert() error
	Delete(id int64) error
}

type inventorySupplierRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.InventorySupplier
}

func (m *inventorySupplierRepo) Create(ii *models.InventorySupplier) error {
	// hash the password
	if err := ii.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *inventorySupplierRepo) Update(ii *models.InventorySupplier) error {
	if ii.ID == 0 {
		return fmt.Errorf("error: can't update ii, missing ii ID")
	}
	if err := ii.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *inventorySupplierRepo) Save() error {
	return m.InventorySupplier.Save(m.ctx, m.db)
}

func (m *inventorySupplierRepo) Upsert() error {
	return m.InventorySupplier.Upsert(m.ctx, m.db)
}

func (m *inventorySupplierRepo) Delete(id int64) error {
	ii, err := models.ProductByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return ii.Delete(m.ctx, m.db)
}

func (m *inventorySupplierRepo) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (m *inventorySupplierRepo) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
