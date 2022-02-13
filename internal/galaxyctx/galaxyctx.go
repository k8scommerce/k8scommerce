package galaxyctx

import "context"

type ctxKey string

const (
	storeIdKey      ctxKey = "storeId"
	skuKey          ctxKey = "skuKey"
	slugKey         ctxKey = "slugKey"
	categoryIdKey   ctxKey = "categoryId"
	categorySlugKey ctxKey = "categorySlug"
	currentPageKey  ctxKey = "currentPage"
	filterKey       ctxKey = "filter"
	pageSizeKey     ctxKey = "pageSize"
	sortOnKey       ctxKey = "sortOn"
)

// store id
func SetStoreId(ctx context.Context, storeId int64) context.Context {
	return context.WithValue(ctx, storeIdKey, storeId)
}

func GetStoreId(ctx context.Context) int64 {
	return ctx.Value(storeIdKey).(int64)
}

// sku
func SetSku(ctx context.Context, sku string) context.Context {
	return context.WithValue(ctx, skuKey, sku)
}

func GetSku(ctx context.Context) string {
	return ctx.Value(skuKey).(string)
}

// slug
func SetSlug(ctx context.Context, slug string) context.Context {
	return context.WithValue(ctx, slugKey, slug)
}

func GetSlug(ctx context.Context) string {
	return ctx.Value(slugKey).(string)
}

// categoryId
func SetCategoryId(ctx context.Context, categoryId int64) context.Context {
	return context.WithValue(ctx, categoryIdKey, categoryId)
}

func GetCategoryId(ctx context.Context) int64 {
	return ctx.Value(categoryIdKey).(int64)
}

// category slug
func SetCategorySlug(ctx context.Context, categorySlug string) context.Context {
	return context.WithValue(ctx, categorySlugKey, categorySlug)
}

func GetCategorySlug(ctx context.Context) string {
	return ctx.Value(categorySlugKey).(string)
}

// current page
func SetCurrentPage(ctx context.Context, currentPage int64) context.Context {
	return context.WithValue(ctx, currentPageKey, currentPage)
}

func GetCurrentPage(ctx context.Context) int64 {
	return ctx.Value(currentPageKey).(int64)
}

// page size
func SetPageSize(ctx context.Context, pageSize int64) context.Context {
	return context.WithValue(ctx, pageSizeKey, pageSize)
}

func GetPageSize(ctx context.Context) int64 {
	return ctx.Value(pageSizeKey).(int64)
}

// filter
func SetFilter(ctx context.Context, filter string) context.Context {
	return context.WithValue(ctx, filterKey, filter)
}

func GetFilter(ctx context.Context) string {
	return ctx.Value(filterKey).(string)
}

// sort on
func SetSortOn(ctx context.Context, sortOn string) context.Context {
	return context.WithValue(ctx, sortOnKey, sortOn)
}

func GetSortOn(ctx context.Context) string {
	return ctx.Value(sortOnKey).(string)
}
