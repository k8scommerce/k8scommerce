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
	GetProductById(productID int64) (res *productResponse, err error)
	GetProductBySku(sku string) (res *productResponse, err error)
	GetProductBySlug(sku string) (res *productResponse, err error)
	GetProductsByCategoryId(categoryID, currentPage, pageSize int64, sortOn string) (
		res *getByCategoryIDResponse,
		err error,
	)
}

type similarProductsRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.Product
}

type productResponse struct {
	Product  models.Product
	Variants []models.Variant
	Prices   []models.Price
}

type getByCategoryIDResults struct {
	Product  models.Product
	Variant  models.Variant
	Category models.Category
	Price    models.Price
}

type getByCategoryIDResponse struct {
	PagingStats PagingStats
	Results     []getByCategoryIDResults
}

// products
func (m *similarProductsRepo) GetProductBySku(sku string) (res *productResponse, err error) {
	nstmt, err := m.db.PrepareNamed(`
		SELECT 
			-- product
			p.id AS "product.id",
			p.slug AS "product.slug",
			p.name AS "product.name",
			p.short_description AS "product.short_description",
			p.description AS "product.description",
			p.meta_title AS "product.meta_title",
			p.meta_description AS "product.meta_description",
			p.meta_keywords AS "product.meta_keywords",
			p.promotionable AS "product.promotionable",
			p.available_on AS "product.available_on",
			p.discontinue_on AS "product.discontinue_on",

			-- variant
			v.id AS "variant.id",
			v.product_id AS "variant.product_id",
			v.is_default AS "variant.is_default",
			v.sku AS "variant.sku",
			v.sort_order AS "variant.sort_order",
			v.cost_amount AS "variant.cost_amount",
			v.cost_currency AS "variant.cost_currency",
			v.track_inventory AS "variant.track_inventory",
			v.tax_category_id AS "variant.tax_category_id",
			v.shipping_category_id AS "variant.shipping_category_id",
			v.discontinue_on AS "variant.discontinue_on",
			v.weight AS "variant.weight",
			v.height AS "variant.height",
			v.width AS "variant.width",
			v.depth AS "variant.depth",

			-- price
			pr.variant_id AS "price.variant_id",
			pr.amount AS "price.amount",
			pr.compare_at_amount AS "price.compare_at_amount",
			pr.currency AS "price.currency",
			pr.user_role_id AS "price.user_role_id"

		FROM product p
		INNER JOIN variant v ON p.id = v.product_id
		INNER JOIN price pr on pr.variant_id = v.id AND pr.user_role_id is null
		WHERE v.sku = :sku;
	`)
	if err != nil {
		return nil, fmt.Errorf("error::GetProductBySku::%s", err.Error())
	}

	var result []*struct {
		Product models.Product `db:"product"`
		Variant models.Variant `db:"variant"`
		Price   models.Price   `db:"price"`
	}
	err = nstmt.Select(&result,
		map[string]interface{}{
			"sku": sku,
		})

	out := &productResponse{}
	if len(result) > 0 {
		out.Product = result[0].Product
		out.Variants = []models.Variant{}
		out.Prices = []models.Price{}
		for _, r := range result {
			out.Variants = append(out.Variants, r.Variant)
			out.Prices = append(out.Prices, r.Price)
		}
		return out, err
	}
	return nil, err
}

func (m *similarProductsRepo) GetProductBySlug(slug string) (res *productResponse, err error) {
	nstmt, err := m.db.PrepareNamed(`
		WITH p AS (
			SELECT 
				product.* 
			FROM product
			WHERE product.slug = :slug
		)
		SELECT 
			-- product
			p.id AS "product.id",
			p.slug AS "product.slug",
			p.name AS "product.name",
			p.short_description AS "product.short_description",
			p.description AS "product.description",
			p.meta_title AS "product.meta_title",
			p.meta_description AS "product.meta_description",
			p.meta_keywords AS "product.meta_keywords",
			p.promotionable AS "product.promotionable",
			p.available_on AS "product.available_on",
			p.discontinue_on AS "product.discontinue_on",

			-- variant
			v.id AS "variant.id",
			v.product_id AS "variant.product_id",
			v.is_default AS "variant.is_default",
			v.sku AS "variant.sku",
			v.sort_order AS "variant.sort_order",
			v.cost_amount AS "variant.cost_amount",
			v.cost_currency AS "variant.cost_currency",
			v.track_inventory AS "variant.track_inventory",
			v.tax_category_id AS "variant.tax_category_id",
			v.shipping_category_id AS "variant.shipping_category_id",
			v.discontinue_on AS "variant.discontinue_on",
			v.weight AS "variant.weight",
			v.height AS "variant.height",
			v.width AS "variant.width",
			v.depth AS "variant.depth",

			-- price
			pr.variant_id AS "price.variant_id",
			pr.amount AS "price.amount",
			pr.compare_at_amount AS "price.compare_at_amount",
			pr.currency AS "price.currency",
			pr.user_role_id AS "price.user_role_id"

		FROM p
		INNER JOIN variant v ON p.id = v.product_id
		INNER JOIN price pr on pr.variant_id = v.id AND pr.user_role_id is null;
	`)
	if err != nil {
		return nil, fmt.Errorf("error::GetProductBySlug::%s", err.Error())
	}

	var result []*struct {
		Product models.Product `db:"product"`
		Variant models.Variant `db:"variant"`
		Price   models.Price   `db:"price"`
	}
	err = nstmt.Select(&result,
		map[string]interface{}{
			"slug": slug,
		})

	out := &productResponse{}
	if len(result) > 0 {
		out.Product = result[0].Product
		out.Variants = []models.Variant{}
		out.Prices = []models.Price{}
		for _, r := range result {
			out.Variants = append(out.Variants, r.Variant)
			out.Prices = append(out.Prices, r.Price)
		}
		return out, err
	}
	return nil, err
}

func (m *similarProductsRepo) GetProductById(productID int64) (res *productResponse, err error) {
	nstmt, err := m.db.PrepareNamed(`
		WITH p AS (
			SELECT 
				product.* 
			FROM product
			WHERE product.id = :product_id
		)
		SELECT 
			-- product
			p.id AS "product.id",
			p.slug AS "product.slug",
			p.name AS "product.name",
			p.short_description AS "product.short_description",
			p.description AS "product.description",
			p.meta_title AS "product.meta_title",
			p.meta_description AS "product.meta_description",
			p.meta_keywords AS "product.meta_keywords",
			p.promotionable AS "product.promotionable",
			p.available_on AS "product.available_on",
			p.discontinue_on AS "product.discontinue_on",

			-- variant
			v.id AS "variant.id",
			v.product_id AS "variant.product_id",
			v.is_default AS "variant.is_default",
			v.sku AS "variant.sku",
			v.sort_order AS "variant.sort_order",
			v.cost_amount AS "variant.cost_amount",
			v.cost_currency AS "variant.cost_currency",
			v.track_inventory AS "variant.track_inventory",
			v.tax_category_id AS "variant.tax_category_id",
			v.shipping_category_id AS "variant.shipping_category_id",
			v.discontinue_on AS "variant.discontinue_on",
			v.weight AS "variant.weight",
			v.height AS "variant.height",
			v.width AS "variant.width",
			v.depth AS "variant.depth",

			-- price
			pr.variant_id AS "price.variant_id",
			pr.amount AS "price.amount",
			pr.compare_at_amount AS "price.compare_at_amount",
			pr.currency AS "price.currency",
			pr.user_role_id AS "price.user_role_id"

		FROM p
		INNER JOIN variant v ON p.id = v.product_id
		INNER JOIN price pr on pr.variant_id = v.id AND pr.user_role_id is null;
	`)
	if err != nil {
		return nil, fmt.Errorf("error::GetProductById::%s", err.Error())
	}

	var result []*struct {
		Product models.Product `db:"product"`
		Variant models.Variant `db:"variant"`
		Price   models.Price   `db:"price"`
	}
	err = nstmt.Select(&res,
		map[string]interface{}{
			"product_id": productID,
		})

	out := &productResponse{}
	if len(result) > 0 {
		out.Product = result[0].Product
		out.Variants = []models.Variant{}
		out.Prices = []models.Price{}
		for _, r := range result {
			out.Variants = append(out.Variants, r.Variant)
			out.Prices = append(out.Prices, r.Price)
		}
		return out, err
	}
	return nil, err
}

func (m *similarProductsRepo) GetProductsByCategoryId(categoryID, currentPage, pageSize int64, sortOn string) (res *getByCategoryIDResponse, err error) {
	fmt.Println("categoryID", categoryID)
	fmt.Println("currentPage", currentPage)
	fmt.Println("pageSize", pageSize)

	orderBy, err := BuildOrderBy(sortOn, map[string]string{
		"name":   "p",  // product alias
		"amount": "pr", // price alias
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
			-- product
			p.id AS "product.id",
			p.slug AS "product.slug",
			p.name AS "product.name",
			p.short_description AS "product.short_description",
			p.description AS "product.description",
			p.meta_title AS "product.meta_title",
			p.meta_description AS "product.meta_description",
			p.meta_keywords AS "product.meta_keywords",
			p.promotionable AS "product.promotionable",
			p.available_on AS "product.available_on",
			p.discontinue_on AS "product.discontinue_on",

			-- variant
			v.id AS "variant.id",
			v.product_id AS "variant.product_id",
			v.is_default AS "variant.is_default",
			v.sku AS "variant.sku",
			v.sort_order AS "variant.sort_order",
			v.cost_amount AS "variant.cost_amount",
			v.cost_currency AS "variant.cost_currency",
			v.track_inventory AS "variant.track_inventory",
			v.tax_category_id AS "variant.tax_category_id",
			v.shipping_category_id AS "variant.shipping_category_id",
			v.discontinue_on AS "variant.discontinue_on",
			v.weight AS "variant.weight",
			v.height AS "variant.height",
			v.width AS "variant.width",
			v.depth AS "variant.depth",

			-- price
			pr.variant_id AS "price.variant_id",
			pr.amount AS "price.amount",
			pr.compare_at_amount AS "price.compare_at_amount",
			pr.currency AS "price.currency",
			pr.user_role_id AS "price.user_role_id",

			-- catgory
			c.id AS "category.id",
			c.parent_id AS "category.parent_id",
			c.store_id AS "category.store_id",
			c.name AS "category.name",
			c.description AS "category.description",
			c.permalink AS "category.permalink",
			c.meta_title AS "category.meta_title",
			c.meta_description AS "category.meta_description",
			c.meta_keywords AS "category.meta_keywords",
			c.hide_from_nav AS "category.hide_from_nav",
			c.lft AS "category.lft",
			c.rgt AS "category.rgt",
			c.depth AS "category.depth",
			c.sort_order AS "category.sort_order",

			-- stats
			COUNT(p.*) OVER() AS "pagingstats.total_records"

		from product p
		inner join product_category pc on p.id = pc.product_id
		inner join category c ON pc.category_id = c.id
		inner join variant v on v.product_id = p.id AND v.is_default = true
		inner join price pr on pr.variant_id = v.id AND pr.user_role_id is null
		WHERE pc.category_id = :category_id
		%s
		%s
		%s
	`, orderBy, offset, limit))
	if err != nil {
		return nil, fmt.Errorf("error::GetProductsByCategoryId::%s", err.Error())
	}

	var result []*struct {
		Product     models.Product
		Variant     models.Variant
		Category    models.Category
		Price       models.Price
		PagingStats PagingStats
	}

	err = nstmt.Select(&result,
		map[string]interface{}{
			"category_id": categoryID,
			"offset":      (currentPage - 1) * pageSize,
			"limit":       pageSize,
			"order_by":    orderBy,
		})

	results := []getByCategoryIDResults{}

	if len(result) > 0 {
		var stats *PagingStats
		for i, r := range result {
			if i == 0 {
				stats = r.PagingStats.Calc(pageSize)
				// totalPages := float64(stats.TotalRecords) / float64(pageSize)
				// stats.TotalPages = int64(math.Ceil(totalPages))
			}
			results = append(results, getByCategoryIDResults{
				Product:  r.Product,
				Variant:  r.Variant,
				Price:    r.Price,
				Category: r.Category,
			})
		}

		out := &getByCategoryIDResponse{
			Results:     results,
			PagingStats: *stats,
		}

		return out, err
	}
	return nil, err
}

func (m *similarProductsRepo) GetAllProducts(currentPage, pageSize int64) ([]*similarProductsRepo, error) {
	return nil, nil
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
