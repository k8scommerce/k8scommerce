package repos

import (
	"context"
	"fmt"

	"github.com/k8scommerce/k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func newInventoryBrand(repo *repo) InventoryBrand {
	return &inventoryBrandRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type InventoryBrand interface {
	Exists() bool
	Deleted() bool
	Create(ii *models.InventoryBrand) error
	Update(ii *models.InventoryBrand) error
	Save() error
	Upsert() error
	Delete(id int64) error
}

type inventoryBrandRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.InventoryBrand
}

func (m *inventoryBrandRepo) Create(ii *models.InventoryBrand) error {
	// hash the password
	if err := ii.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *inventoryBrandRepo) Update(ii *models.InventoryBrand) error {
	if ii.ID == 0 {
		return fmt.Errorf("error: can't update ii, missing ii ID")
	}
	if err := ii.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *inventoryBrandRepo) Save() error {
	return m.InventoryBrand.Save(m.ctx, m.db)
}

func (m *inventoryBrandRepo) Upsert() error {
	return m.InventoryBrand.Upsert(m.ctx, m.db)
}

func (m *inventoryBrandRepo) Delete(id int64) error {
	ii, err := models.ProductByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return ii.Delete(m.ctx, m.db)
}

func (m *inventoryBrandRepo) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (m *inventoryBrandRepo) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
