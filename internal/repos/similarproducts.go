package repos

import (
	"context"
	"fmt"

	"k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
)

func newSimilarProducts(repo *repo) SimilarProducts {
	return &similarProductsRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type SimilarProducts interface {
	Exists() bool
	Deleted() bool
	Create(prod *models.Product) error
	Update(prod *models.Product) error
	Save() error
	Upsert() error
	Delete(id int64) error
}

type similarProductsRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.Product
}

func (m *similarProductsRepo) Create(prod *models.Product) error {
	if err := prod.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *similarProductsRepo) Update(prod *models.Product) error {
	if prod.ID == 0 {
		return fmt.Errorf("error: can't update product, missing product ID")
	}
	if err := prod.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *similarProductsRepo) Save() error {
	return m.Product.Save(m.ctx, m.db)
}

func (m *similarProductsRepo) Upsert() error {
	return m.Product.Upsert(m.ctx, m.db)
}

func (m *similarProductsRepo) Delete(id int64) error {
	prod, err := models.ProductByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return prod.Delete(m.ctx, m.db)
}

// func (m *similarProductsRepo) getProductVariants(productId int64) (*[]types.Variant, error) {
// 	nstmt, err := m.db.PrepareNamed(`
// 		SELECT
// 			*
// 		FROM variants
// 		WHERE product_id = :product_id
// 	`)
// 	if err != nil {
// 		return nil, fmt.Errorf("error::GetProductsByCategoryId::%s", err.Error())
// 	}

// 	v := []types.Variant{}
// 	err = nstmt.Select(&v,
// 		map[string]interface{}{
// 			"product_id": productId,
// 		})
// 	if err != nil {
// 		return nil, err
// 	}

// 	b, _ := json.MarshalIndent(&v, "", "    ")
// 	fmt.Printf("results: %s", string(b))

// 	return &v, err
// }
