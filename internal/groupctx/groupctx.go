package groupctx

import (
	"context"
	"k8scommerce/services/rpc/cart/pb/cart"
)

type ctxKey string

const (
	storeIdKey       ctxKey = "storeId"
	categoryIdKey    ctxKey = "categoryId"
	categorySlugKey  ctxKey = "categorySlug"
	cartKey          ctxKey = "cart"
	customerEmailKey ctxKey = "customerEmail"
	productIdKey     ctxKey = "productId"
	productSlugKey   ctxKey = "productSlug"
	productSkuKey    ctxKey = "productSku"
	currentPageKey   ctxKey = "currentPage"
	filterKey        ctxKey = "filter"
	pageSizeKey      ctxKey = "pageSize"
	sortOnKey        ctxKey = "sortOn"
	warehouseIdKey   ctxKey = "warehouseId"
)

// store id
func SetStoreId(ctx context.Context, storeId int64) context.Context {
	return context.WithValue(ctx, storeIdKey, storeId)
}

func GetStoreId(ctx context.Context) int64 {
	if ctx.Value(storeIdKey) != nil {
		return ctx.Value(storeIdKey).(int64)
	}
	return 0
}

// categoryId
func SetCategoryId(ctx context.Context, categoryId int64) context.Context {
	return context.WithValue(ctx, categoryIdKey, categoryId)
}

func GetCategoryId(ctx context.Context) int64 {
	if ctx.Value(categoryIdKey) != nil {
		return ctx.Value(categoryIdKey).(int64)
	}
	return 0
}

// category slug
func SetCategorySlug(ctx context.Context, categorySlug string) context.Context {
	return context.WithValue(ctx, categorySlugKey, categorySlug)
}

func GetCategorySlug(ctx context.Context) string {
	if ctx.Value(categorySlugKey) != nil {
		return ctx.Value(categorySlugKey).(string)
	}
	return ""
}

// category slug
func SetCustomerEmail(ctx context.Context, customerEmail string) context.Context {
	return context.WithValue(ctx, customerEmailKey, customerEmail)
}

func GetCustomerEmail(ctx context.Context) string {
	if ctx.Value(customerEmailKey) != nil {
		return ctx.Value(customerEmailKey).(string)
	}
	return ""
}

// productId
func SetProductId(ctx context.Context, productId int64) context.Context {
	return context.WithValue(ctx, productIdKey, productId)
}

func GetProductId(ctx context.Context) int64 {
	if ctx.Value(productIdKey) != nil {
		return ctx.Value(productIdKey).(int64)
	}
	return 0
}

// product slug
func SetProductSlug(ctx context.Context, productSlug string) context.Context {
	return context.WithValue(ctx, productSlugKey, productSlug)
}

func GetProductSlug(ctx context.Context) string {
	if ctx.Value(productSlugKey) != nil {
		return ctx.Value(productSlugKey).(string)
	}
	return ""
}

// product slug
func SetProductSku(ctx context.Context, productSku string) context.Context {
	return context.WithValue(ctx, productSkuKey, productSku)
}

func GetProductSku(ctx context.Context) string {
	if ctx.Value(productSkuKey) != nil {
		return ctx.Value(productSkuKey).(string)
	}
	return ""
}

// current page
func SetCurrentPage(ctx context.Context, currentPage int64) context.Context {
	return context.WithValue(ctx, currentPageKey, currentPage)
}

func GetCurrentPage(ctx context.Context) int64 {
	if ctx.Value(currentPageKey) != nil {
		return ctx.Value(currentPageKey).(int64)
	}
	return 0
}

// page size
func SetPageSize(ctx context.Context, pageSize int64) context.Context {
	return context.WithValue(ctx, pageSizeKey, pageSize)
}

func GetPageSize(ctx context.Context) int64 {
	if ctx.Value(pageSizeKey) != nil {
		return ctx.Value(pageSizeKey).(int64)
	}
	return 0
}

// filter
func SetFilter(ctx context.Context, filter string) context.Context {
	return context.WithValue(ctx, filterKey, filter)
}

func GetFilter(ctx context.Context) string {
	if ctx.Value(filterKey) != nil {
		return ctx.Value(filterKey).(string)
	}
	return ""
}

// sort on
func SetSortOn(ctx context.Context, sortOn string) context.Context {
	return context.WithValue(ctx, sortOnKey, sortOn)
}

func GetSortOn(ctx context.Context) string {
	if ctx.Value(sortOnKey) != nil {
		return ctx.Value(sortOnKey).(string)
	}
	return ""
}

// warehouse id
func SetWarehouseId(ctx context.Context, warehouseId int64) context.Context {
	return context.WithValue(ctx, warehouseIdKey, warehouseId)
}

func GetWarehouseId(ctx context.Context) int64 {
	if ctx.Value(warehouseIdKey) != nil {
		return ctx.Value(warehouseIdKey).(int64)
	}
	return 0

}

// store id
func SetCart(ctx context.Context, ct *cart.Cart) context.Context {
	return context.WithValue(ctx, cartKey, ct)
}

func GetCart(ctx context.Context) *cart.Cart {
	if ctx.Value(cartKey) != nil {
		return ctx.Value(cartKey).(*cart.Cart)
	}
	return nil
}
