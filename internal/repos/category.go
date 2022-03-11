package repos

import (
	"context"
	"fmt"

	"k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
)

func newCategory(repo *repo) Category {
	return &categoryRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type Category interface {
	Exists() bool
	Deleted() bool
	Create(cat *models.Category) error
	Update(cat *models.Category) error
	Save() error
	Upsert() error
	Delete(id int64) error
	GetCategoryBySlug(storeId int64, slug string) (*models.Category, error)
	GetCategoryById(id int64) (*models.Category, error)
	GetAllCategories(storeId int64) (res *getAllCategoriesResponse, err error)
}

type categoryRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.Category
}

type getAllCategoriesResponse struct {
	PagingStats PagingStats
	Categories  []models.Category
}

func (m *categoryRepo) GetCategoryBySlug(storeId int64, slug string) (*models.Category, error) {
	return models.CategoryByStoreIDSlug(m.ctx, m.db, storeId, slug)
}

func (m *categoryRepo) GetCategoryById(id int64) (*models.Category, error) {
	return models.CategoryByID(m.ctx, m.db, id)
}

func (m *categoryRepo) GetAllCategories(storeId int64) (res *getAllCategoriesResponse, err error) {
	nstmt, err := m.db.PrepareNamed(`
			select 
				-- catgory
				c.id AS "category.id",
				c.parent_id AS "category.parent_id",
				c.slug AS "category.slug",
				c.name AS "category.name",
				c.description AS "category.description",
				c.meta_title AS "category.meta_title",
				c.meta_description AS "category.meta_description",
				c.meta_keywords AS "category.meta_keywords",
				c.hide_from_nav AS "category.hide_from_nav",
				c.lft AS "category.lft",
				c.rgt AS "category.rgt",
				c.depth AS "category.depth",
				c.sort_order AS "category.sort_order"
			from category c
			where c.store_id = :store_id
			ORDER BY c.lft ASC
		`)
	if err != nil {
		return nil, fmt.Errorf("error::GetAllCategories::%s", err.Error())
	}

	var result []*struct {
		Category models.Category
	}

	err = nstmt.Select(&result,
		map[string]interface{}{
			"store_id": storeId,
		})

	var categories []models.Category
	if len(result) > 0 {
		for _, r := range result {
			categories = append(categories, r.Category)
		}

		out := &getAllCategoriesResponse{
			Categories: categories,
		}
		return out, err
	}

	return nil, err
}

func (m *categoryRepo) Create(cat *models.Category) error {
	if err := cat.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *categoryRepo) Update(cat *models.Category) error {
	if cat.ID == 0 {
		return fmt.Errorf("error: can't update category, missing category ID")
	}
	if err := cat.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *categoryRepo) Save() error {
	return m.Category.Save(m.ctx, m.db)
}

func (m *categoryRepo) Upsert() error {
	return m.Category.Upsert(m.ctx, m.db)
}

func (m *categoryRepo) Delete(id int64) error {
	cat, err := models.CategoryByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return cat.Delete(m.ctx, m.db)
}
