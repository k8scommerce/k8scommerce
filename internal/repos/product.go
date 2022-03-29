package repos

import (
	"context"
	"fmt"

	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/k8scommerce/k8scommerce/internal/buildsql"
	"github.com/k8scommerce/k8scommerce/internal/models"

	"github.com/jmoiron/sqlx"
)

func newProduct(repo *repo) Product {
	return &productRepo{
		db:   repo.db,
		repo: repo,
		ctx:  context.Background(),
	}
}

type Product interface {
	Exists() bool
	Deleted() bool
	Create(prod *models.Product) error
	Update(prod *models.Product) error
	Save() error
	Upsert() error
	Delete(id int64) error
	GetProductById(productID int64) (res *productResponse, err error)
	GetProductBySku(storeId int64, sku string) (res *productResponse, err error)
	GetProductBySlug(storeId int64, slug string) (res *productResponse, err error)
	GetProductsByCategoryId(storeId, categoryID, currentPage, pageSize int64, filter string) (
		res *getProductsByCategoryResponse,
		err error,
	)
	GetProductsByCategorySlug(storeId int64, categorySlug string, currentPage, pageSize int64, filter string) (
		res *getProductsByCategoryResponse,
		err error,
	)
	GetAllProducts(storeId, currentPage, pageSize int64, filter string) (
		res *getAllProductsResponse,
		err error,
	)
}

type productRepo struct {
	db   *sqlx.DB
	repo *repo
	ctx  context.Context

	*models.Product
}
type productResponse struct {
	Product    models.Product
	Variants   []models.Variant
	Prices     []models.Price
	Categories []CategoryPair
}

type getAllProductsResults struct {
	Product models.Product
	Variant models.Variant
	Price   models.Price
	Asset   models.Asset
}

type getProductsByCategoryResponse struct {
	PagingStats PagingStats
	Results     []getAllProductsResults
}

type getAllProductsResponse struct {
	PagingStats PagingStats
	Results     []getAllProductsResults
}

// products
func (m *productRepo) GetProductBySku(storeId int64, sku string) (res *productResponse, err error) {
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
			pr.sale_price AS "price.sale_price",
			pr.retail_price AS "price.retail_price",
			pr.currency AS "price.currency",

			-- categories
			array_to_json(c.categories) AS "categories"

		FROM product p
		INNER JOIN variant v ON p.id = v.product_id
		INNER JOIN price pr ON pr.variant_id = v.id
		CROSS JOIN LATERAL (
			SELECT ARRAY (
				select 
					json_build_object(
						'slug', c.slug,
						'name', c.name
					)
				from product_category pc
				inner join category c ON pc.category_id = c.id
				where p.id = pc.product_id
			) as categories
		) c
		WHERE v.sku = :sku 
		AND p.store_id = :store_id;
	`)
	if err != nil {
		return nil, fmt.Errorf("error::GetProductBySku::%s", err.Error())
	}

	var result []*struct {
		Product    models.Product `db:"product"`
		Variant    models.Variant `db:"variant"`
		Price      models.Price   `db:"price"`
		Categories CategoryPairs  `db:"categories"`
	}
	err = nstmt.Select(&result,
		map[string]interface{}{
			"store_id": storeId,
			"sku":      sku,
		})

	out := &productResponse{}
	if len(result) > 0 {
		out.Product = result[0].Product
		out.Variants = []models.Variant{}
		out.Prices = []models.Price{}
		out.Categories = result[0].Categories.Pairs
		for _, r := range result {
			out.Variants = append(out.Variants, r.Variant)
			out.Prices = append(out.Prices, r.Price)
		}
		return out, err
	}
	return nil, err
}

func (m *productRepo) GetProductBySlug(storeId int64, slug string) (res *productResponse, err error) {
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
			pr.sale_price AS "price.sale_price",
			pr.retail_price AS "price.retail_price",
			pr.currency AS "price.currency",

			-- categories
			array_to_json(c.categories) AS "categories"

		FROM p
		INNER JOIN variant v ON p.id = v.product_id
		INNER JOIN price pr ON pr.variant_id = v.id
		CROSS JOIN LATERAL (
			SELECT ARRAY (
				select 
					json_build_object(
						'slug', c.slug,
						'name', c.name
					)
				from product_category pc
				inner join category c ON pc.category_id = c.id
				where p.id = pc.product_id
			) as categories
		) c
		WHERE p.store_id = :store_id;
	`)
	if err != nil {
		return nil, fmt.Errorf("error::GetProductBySlug::%s", err.Error())
	}

	var result []*struct {
		Product    models.Product `db:"product"`
		Variant    models.Variant `db:"variant"`
		Price      models.Price   `db:"price"`
		Categories CategoryPairs  `db:"categories"`
	}

	err = nstmt.Select(&result,
		map[string]interface{}{
			"store_id": storeId,
			"slug":     slug,
		})

	out := &productResponse{}
	if len(result) > 0 {
		out.Product = result[0].Product
		out.Variants = []models.Variant{}
		out.Prices = []models.Price{}
		out.Categories = result[0].Categories.Pairs
		for _, r := range result {
			out.Variants = append(out.Variants, r.Variant)
			out.Prices = append(out.Prices, r.Price)
		}
		return out, err
	}
	return nil, err
}

func (m *productRepo) GetProductById(productID int64) (res *productResponse, err error) {
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
			pr.sale_price AS "price.sale_price",
			pr.retail_price AS "price.retail_price",
			pr.currency AS "price.currency",
	
			-- categories
			array_to_json(c.categories) AS "categories"

		FROM p
		INNER JOIN variant v ON p.id = v.product_id
		INNER JOIN price pr ON pr.variant_id = v.id 
		CROSS JOIN LATERAL (
			SELECT ARRAY (
				select 
					json_build_object(
						'slug', c.slug,
						'name', c.name
					)
				from product_category pc
				inner join category c ON pc.category_id = c.id
				where p.id = pc.product_id
			) as categories
		) c;
	`)
	if err != nil {
		return nil, fmt.Errorf("error::GetProductById::%s", err.Error())
	}

	var result []*struct {
		Product    models.Product `db:"product"`
		Variant    models.Variant `db:"variant"`
		Price      models.Price   `db:"price"`
		Categories CategoryPairs  `db:"categories"`
	}
	err = nstmt.Select(&result,
		map[string]interface{}{
			"product_id": productID,
		})

	out := &productResponse{}
	if len(result) > 0 {
		out.Product = result[0].Product
		out.Variants = []models.Variant{}
		out.Prices = []models.Price{}
		out.Categories = result[0].Categories.Pairs
		for _, r := range result {
			out.Variants = append(out.Variants, r.Variant)
			out.Prices = append(out.Prices, r.Price)
		}
		return out, err
	}
	return nil, err
}

func (m *productRepo) GetProductsByCategoryId(storeId, categoryId, currentPage, pageSize int64, filter string) (res *getProductsByCategoryResponse, err error) {

	var builder = buildsql.NewQueryBuilder()
	where, orderBy, namedParamMap, err := builder.Build(filter, map[string]interface{}{
		"p":  models.Product{}, // product alias
		"v":  models.Variant{}, // product alias
		"pr": models.Price{},   // product alias
	})
	if err != nil {
		return nil, err
	}

	// set a default order by
	if orderBy == "" {
		orderBy = "ORDER BY p.name ASC"
	}
	offset := fmt.Sprintf("OFFSET %d", currentPage*pageSize)
	limit := fmt.Sprintf("LIMIT %d", pageSize)

	sql := fmt.Sprintf(`
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
			pr.sale_price AS "price.sale_price",
			pr.retail_price AS "price.retail_price",
			pr.currency AS "price.currency",

			-- asset
			-- COALESCE(a.url,'') AS "asset.url", -- we don't need to give the fill size image away
			COALESCE(a.display_name,'') AS "asset.display_name",
			COALESCE(a.sizes,'{}') AS "asset.sizes",

			-- stats
			COUNT(p.*) OVER() AS "pagingstats.total_records"

		from product p
		inner join product_category pc on p.id = pc.product_id
		inner join category c ON pc.category_id = c.id
		inner join variant v on v.product_id = p.id AND v.is_default = true
		inner join price pr ON pr.variant_id = v.id
		LEFT JOIN LATERAL (
			SELECT * FROM asset
			WHERE asset.product_id = p.id
			AND asset.variant_id = v.id
			AND asset.kind = :asset_kind
			ORDER BY asset.sort_order ASC
			LIMIT 1
		) a ON a.product_id = p.id 
		where p.store_id = :store_id
		and c.id = :category_id
		%s
		%s
		%s
		%s
	`, where, orderBy, offset, limit)

	fmt.Println(sql)

	nstmt, err := m.db.PrepareNamed(sql)
	if err != nil {

		// fmt.Println("where", where)
		// fmt.Println("orderBy", orderBy)
		// fmt.Println("offset", offset)
		// fmt.Println("limit", limit)

		return nil, fmt.Errorf("error::GetProductsByCategoryId::%s", err.Error())
	}

	var result []*struct {
		Product     models.Product
		Variant     models.Variant
		Price       models.Price
		Asset       models.Asset
		PagingStats PagingStats
	}

	namedParamMap["store_id"] = storeId
	namedParamMap["category_id"] = categoryId
	namedParamMap["offset"] = currentPage * pageSize
	namedParamMap["limit"] = pageSize

	// intentionally hardcode the asset kind
	namedParamMap["asset_kind"] = catalog.AssetKind_image.Number()

	err = nstmt.Select(&result, namedParamMap)

	results := []getAllProductsResults{}

	if len(result) > 0 {
		var stats *PagingStats
		for i, r := range result {
			if i == 0 {
				stats = r.PagingStats.Calc(pageSize)
			}

			results = append(results, getAllProductsResults{
				Product: r.Product,
				Variant: r.Variant,
				Price:   r.Price,
				Asset:   r.Asset,
			})
		}

		out := &getProductsByCategoryResponse{
			Results:     results,
			PagingStats: *stats,
		}

		return out, err
	}
	return nil, err
}

func (m *productRepo) GetProductsByCategorySlug(storeId int64, categorySlug string, currentPage, pageSize int64, filter string) (res *getProductsByCategoryResponse, err error) {

	var builder = buildsql.NewQueryBuilder()
	where, orderBy, namedParamMap, err := builder.Build(filter, map[string]interface{}{
		"p":  models.Product{}, // product alias
		"v":  models.Variant{}, // product alias
		"pr": models.Price{},   // product alias
	})
	if err != nil {
		return nil, err
	}

	// set a default order by
	if orderBy == "" {
		orderBy = "ORDER BY p.name ASC"
	}
	offset := fmt.Sprintf("OFFSET %d", currentPage*pageSize)
	limit := fmt.Sprintf("LIMIT %d", pageSize)

	sql := fmt.Sprintf(`
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
			pr.sale_price AS "price.retail_price",
			pr.retail_price AS "price.compare_at_amount",
			pr.currency AS "price.currency",

			-- asset
			-- COALESCE(a.url,'') AS "asset.url", -- we don't need to give the fill size image away
			COALESCE(a.display_name,'') AS "asset.display_name",
			COALESCE(a.sizes,'{}') AS "asset.sizes",

			-- stats
			COUNT(p.*) OVER() AS "pagingstats.total_records"

		from product p
		inner join product_category pc on p.id = pc.product_id
		inner join category c ON pc.category_id = c.id
		inner join variant v on v.product_id = p.id AND v.is_default = true
		inner join price pr ON pr.variant_id = v.id
		LEFT JOIN LATERAL (
			SELECT * FROM asset
			WHERE asset.product_id = p.id
			AND asset.variant_id = v.id
			AND asset.kind = :asset_kind
			ORDER BY asset.sort_order ASC
			LIMIT 1
		) a ON a.product_id = p.id 
		where p.store_id = :store_id
		%s
		%s
		%s
		%s
	`, where, orderBy, offset, limit)

	fmt.Println(sql)

	nstmt, err := m.db.PrepareNamed(sql)
	if err != nil {

		return nil, fmt.Errorf("error::GetAllProducts::%s", err.Error())
	}

	var result []*struct {
		Product     models.Product
		Variant     models.Variant
		Price       models.Price
		Asset       models.Asset
		PagingStats PagingStats
	}

	namedParamMap["store_id"] = storeId
	namedParamMap["offset"] = currentPage * pageSize
	namedParamMap["limit"] = pageSize

	// intentionally hardcode the asset kind
	namedParamMap["asset_kind"] = catalog.AssetKind_image.Number()

	err = nstmt.Select(&result, namedParamMap)

	results := []getAllProductsResults{}

	if len(result) > 0 {
		var stats *PagingStats
		for i, r := range result {
			if i == 0 {
				stats = r.PagingStats.Calc(pageSize)
			}

			results = append(results, getAllProductsResults{
				Product: r.Product,
				Variant: r.Variant,
				Price:   r.Price,
				Asset:   r.Asset,
			})
		}

		out := &getProductsByCategoryResponse{
			Results:     results,
			PagingStats: *stats,
		}

		return out, err
	}
	return nil, err
}

func (m *productRepo) GetAllProducts(storeId, currentPage, pageSize int64, filter string) (res *getAllProductsResponse, err error) {

	var builder = buildsql.NewQueryBuilder()
	where, orderBy, namedParamMap, err := builder.Build(filter, map[string]interface{}{
		"p":  models.Product{}, // product alias
		"v":  models.Variant{}, // product alias
		"pr": models.Price{},   // product alias
	})
	if err != nil {
		return nil, err
	}

	// set a default order by
	if orderBy == "" {
		orderBy = "ORDER BY p.name ASC"
	}
	offset := fmt.Sprintf("OFFSET %d", currentPage*pageSize)
	limit := fmt.Sprintf("LIMIT %d", pageSize)

	sql := fmt.Sprintf(`
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
			pr.sale_price AS "price.sale_price",
			pr.retail_price AS "price.retail_price",
			pr.currency AS "price.currency",

			-- asset
			-- COALESCE(a.url,'') AS "asset.url", -- we don't need to give the fill size image away
			COALESCE(a.display_name,'') AS "asset.display_name",
			COALESCE(a.sizes,'{}') AS "asset.sizes",

			-- stats
			COUNT(p.*) OVER() AS "pagingstats.total_records"

		from product p
		inner join variant v on v.product_id = p.id AND v.is_default = true
		inner join price pr ON pr.variant_id = v.id
		LEFT JOIN LATERAL (
			SELECT display_name, sizes FROM asset
			WHERE asset.variant_id = v.id
			AND asset.kind = 1
			ORDER BY asset.sort_order ASC
			LIMIT 1
   		) a on 1 = 1
		where p.store_id = :store_id
		%s
		%s
		%s
		%s
	`, where, orderBy, offset, limit)

	nstmt, err := m.db.PrepareNamed(sql)
	if err != nil {

		// fmt.Println("where", where)
		// fmt.Println("orderBy", orderBy)
		// fmt.Println("offset", offset)
		// fmt.Println("limit", limit)

		return nil, fmt.Errorf("error::GetAllProducts::%s", err.Error())
	}

	var result []*struct {
		Product     models.Product
		Variant     models.Variant
		Price       models.Price
		Asset       models.Asset
		PagingStats PagingStats
	}

	namedParamMap["store_id"] = storeId
	namedParamMap["offset"] = currentPage * pageSize
	namedParamMap["limit"] = pageSize

	// intentionally hardcode the asset kind
	namedParamMap["asset_kind"] = catalog.AssetKind_image.Number()

	err = nstmt.Select(&result, namedParamMap)

	results := []getAllProductsResults{}

	if len(result) > 0 {
		var stats *PagingStats
		for i, r := range result {
			if i == 0 {
				stats = r.PagingStats.Calc(pageSize)
			}

			results = append(results, getAllProductsResults{
				Product: r.Product,
				Variant: r.Variant,
				Price:   r.Price,
				Asset:   r.Asset,
			})
		}

		out := &getAllProductsResponse{
			Results:     results,
			PagingStats: *stats,
		}

		return out, err
	}
	return nil, err
}

func (m *productRepo) Create(prod *models.Product) error {
	if err := prod.Insert(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *productRepo) Update(prod *models.Product) error {
	if prod.ID == 0 {
		return fmt.Errorf("error: can't update product, missing product ID")
	}
	if err := prod.Update(m.ctx, m.db); err != nil {
		return err
	}
	return nil
}

func (m *productRepo) Save() error {
	return m.Product.Save(m.ctx, m.db)
}

func (m *productRepo) Upsert() error {
	return m.Product.Upsert(m.ctx, m.db)
}

func (m *productRepo) Delete(id int64) error {
	prod, err := models.ProductByID(m.ctx, m.db, id)
	if err != nil {
		return err
	}
	return prod.Delete(m.ctx, m.db)
}

// func (m *productRepo) getProductVariants(productId int64) (*[]types.Variant, error) {
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
