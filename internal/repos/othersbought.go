package repos

import (
	"context"
	"fmt"

	"k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
)

func newOthersBought(repo *repo) OthersBought {
	return &othersBoughtRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

//go:generate mockgen -destination=./mocks/OthersBought.go -package=mock_repos k8scommerce/internal/repos OthersBought
type OthersBought interface {
	Exists() bool
	Deleted() bool
	Create(prod *models.Product) error
	Update(prod *models.Product) error
	Save() error
	Upsert() error
	Delete(id int64) error
}

type othersBoughtRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.Product
}

func (m *othersBoughtRepo) Create(prod *models.Product) error {
	if err := prod.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *othersBoughtRepo) Update(prod *models.Product) error {
	if prod.ID == 0 {
		return fmt.Errorf("error: can't update product, missing product ID")
	}
	if err := prod.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *othersBoughtRepo) Save() error {
	return m.Product.Save(m.ctx, m.db)
}

func (m *othersBoughtRepo) Upsert() error {
	return m.Product.Upsert(m.ctx, m.db)
}

func (m *othersBoughtRepo) Delete(id int64) error {
	prod, err := models.ProductByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return prod.Delete(m.ctx, m.db)
}
