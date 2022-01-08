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
	GetAllCategories(storeId int64, currentPage, pageSize int64, sortOn string) (res *getAllCategoriesResponse, err error)
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

func (m *categoryRepo) GetAllCategories(storeId int64, currentPage, pageSize int64, sortOn string) (res *getAllCategoriesResponse, err error) {
	fmt.Println("currentPage", currentPage)
	fmt.Println("pageSize", pageSize)

	orderBy, err := BuildOrderBy(sortOn, map[string]string{
		"parent_id":  "s",
		"sort_order": "s",
		"rgt":        "s",
		"lft":        "s",
		"name":       "s",
		"depth":      "s",
	})
	if err != nil {
		return nil, err
	}

	// set a default order by
	if orderBy == "" {
		orderBy = "ORDER BY p.name ASC"
	}
	offset := fmt.Sprintf("OFFSET %d", (currentPage-1)*pageSize)
	limit := fmt.Sprintf("LIMIT %d", pageSize)

	nstmt, err := m.db.PrepareNamed(fmt.Sprintf(`
			select 
				-- catgory
				c.id,
				c.parent_id,
				c.name,
				c.description,
				c.permalink,
				c.meta_title,
				c.meta_description,
				c.meta_keywords,
				c.hide_from_nav,
				c.lft,
				c.rgt,
				c.depth,
				c.sort_order,
				
				-- stats
				COUNT(c.*) OVER() AS "pagingstats.total_records"
			from category c 
			where store_id = :store_id
			%s
			%s
			%s
		`, orderBy, offset, limit))
	if err != nil {
		return nil, fmt.Errorf("error::GetAllCategories::%s", err.Error())
	}

	var result *struct {
		Categories  []models.Category
		PagingStats PagingStats
	}

	err = nstmt.Select(&result,
		map[string]interface{}{
			"store_id": storeId,
			"offset":   (currentPage - 1) * pageSize,
			"limit":    pageSize,
			"order_by": orderBy,
		})

	if result != nil && len(result.Categories) > 0 {
		var stats *PagingStats
		stats = result.PagingStats.Calc(pageSize)

		out := &getAllCategoriesResponse{
			Categories:  result.Categories,
			PagingStats: *stats,
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
		return fmt.Errorf("error: can't update cateogry, missing cateogry ID")
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
