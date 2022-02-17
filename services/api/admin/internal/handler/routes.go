// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	api "k8scommerce/services/api/admin/internal/handler/api"
	assets "k8scommerce/services/api/admin/internal/handler/assets"
	categories "k8scommerce/services/api/admin/internal/handler/categories"
	customers "k8scommerce/services/api/admin/internal/handler/customers"
	products "k8scommerce/services/api/admin/internal/handler/products"
	users "k8scommerce/services/api/admin/internal/handler/users"
	"k8scommerce/services/api/admin/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Locale},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/v1/api/ping",
					Handler: api.PingHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Locale, serverCtx.StoreKey},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/v1/categories/:storeId/:currentPage/:pageSize",
					Handler: categories.GetAllCategoriesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/category/slug/:slug",
					Handler: categories.GetCategoryBySlugHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/category/:id",
					Handler: categories.GetCategoryByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/category",
					Handler: categories.CreateCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/v1/category/:id",
					Handler: categories.UpdateCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/v1/category/:id",
					Handler: categories.DeleteCategoryHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Locale, serverCtx.Filter, serverCtx.StoreKey},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/v1/product/sku/:sku",
					Handler: products.GetProductBySkuHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/product/slug/:slug",
					Handler: products.GetProductBySlugHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/product/:id",
					Handler: products.GetProductByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/products/:categoryId/:currentPage/:pageSize",
					Handler: products.GetProductsByCategoryIdHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/v1/products/:currentPage/:pageSize",
					Handler: products.GetAllProductsHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/product",
					Handler: products.CreateProductHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/v1/product/:id",
					Handler: products.UpdateProductHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/v1/product/:id",
					Handler: products.DeleteProductHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Locale, serverCtx.Filter, serverCtx.StoreKey},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v1/asset/:productId/:variantId",
					Handler: assets.UploadHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Locale, serverCtx.StoreKey},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v1/customer",
					Handler: customers.CreateCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/v1/customer/login",
					Handler: customers.CustomerLoginHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Locale},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/v1/user/login",
					Handler: users.LoginHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Locale},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/v1/users/:currentPage/:pageSize",
					Handler: users.GetAllUsersHandler(serverCtx),
				},
			}...,
		),
	)
}
