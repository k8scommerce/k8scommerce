// Code generated by goctl. DO NOT EDIT!
// Source: catalog.proto

package server

import (
	"context"

	"k8scommerce/services/rpc/catalog/internal/logic"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/localrivet/galaxycache"
)

type CatalogClientServer struct {
	svcCtx   *svc.ServiceContext
	universe *galaxycache.Universe
}

func NewCatalogClientServer(svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *CatalogClientServer {
	return &CatalogClientServer{
		svcCtx:   svcCtx,
		universe: universe,
	}
}

//  categories
func (s *CatalogClientServer) GetAllCategories(ctx context.Context, in *catalog.GetAllCategoriesRequest) (*catalog.GetAllCategoriesResponse, error) {
	l := logic.NewGetAllCategoriesLogic(ctx, s.svcCtx, s.universe)
	return l.GetAllCategories(in)
}

func (s *CatalogClientServer) GetCategoryBySlug(ctx context.Context, in *catalog.GetCategoryBySlugRequest) (*catalog.GetCategoryBySlugResponse, error) {
	l := logic.NewGetCategoryBySlugLogic(ctx, s.svcCtx, s.universe)
	return l.GetCategoryBySlug(in)
}

func (s *CatalogClientServer) GetCategoryById(ctx context.Context, in *catalog.GetCategoryByIdRequest) (*catalog.GetCategoryByIdResponse, error) {
	l := logic.NewGetCategoryByIdLogic(ctx, s.svcCtx, s.universe)
	return l.GetCategoryById(in)
}

func (s *CatalogClientServer) CreateCategory(ctx context.Context, in *catalog.CreateCategoryRequest) (*catalog.CreateCategoryResponse, error) {
	l := logic.NewCreateCategoryLogic(ctx, s.svcCtx, s.universe)
	return l.CreateCategory(in)
}

func (s *CatalogClientServer) UpdateCategory(ctx context.Context, in *catalog.UpdateCategoryRequest) (*catalog.UpdateCategoryResponse, error) {
	l := logic.NewUpdateCategoryLogic(ctx, s.svcCtx, s.universe)
	return l.UpdateCategory(in)
}

func (s *CatalogClientServer) DeleteCategory(ctx context.Context, in *catalog.DeleteCategoryRequest) (*catalog.DeleteCategoryResponse, error) {
	l := logic.NewDeleteCategoryLogic(ctx, s.svcCtx, s.universe)
	return l.DeleteCategory(in)
}

//  products
func (s *CatalogClientServer) GetProductBySku(ctx context.Context, in *catalog.GetProductBySkuRequest) (*catalog.GetProductBySkuResponse, error) {
	l := logic.NewGetProductBySkuLogic(ctx, s.svcCtx, s.universe)
	return l.GetProductBySku(in)
}

func (s *CatalogClientServer) GetProductBySlug(ctx context.Context, in *catalog.GetProductBySlugRequest) (*catalog.GetProductBySlugResponse, error) {
	l := logic.NewGetProductBySlugLogic(ctx, s.svcCtx, s.universe)
	return l.GetProductBySlug(in)
}

func (s *CatalogClientServer) GetProductById(ctx context.Context, in *catalog.GetProductByIdRequest) (*catalog.GetProductByIdResponse, error) {
	l := logic.NewGetProductByIdLogic(ctx, s.svcCtx, s.universe)
	return l.GetProductById(in)
}

func (s *CatalogClientServer) GetProductsByCategoryId(ctx context.Context, in *catalog.GetProductsByCategoryIdRequest) (*catalog.GetProductsByCategoryIdResponse, error) {
	l := logic.NewGetProductsByCategoryIdLogic(ctx, s.svcCtx, s.universe)
	return l.GetProductsByCategoryId(in)
}

func (s *CatalogClientServer) GetProductsByCategorySlug(ctx context.Context, in *catalog.GetProductsByCategorySlugRequest) (*catalog.GetProductsByCategorySlugResponse, error) {
	l := logic.NewGetProductsByCategorySlugLogic(ctx, s.svcCtx, s.universe)
	return l.GetProductsByCategorySlug(in)
}

func (s *CatalogClientServer) GetAllProducts(ctx context.Context, in *catalog.GetAllProductsRequest) (*catalog.GetAllProductsResponse, error) {
	l := logic.NewGetAllProductsLogic(ctx, s.svcCtx, s.universe)
	return l.GetAllProducts(in)
}

func (s *CatalogClientServer) CreateProduct(ctx context.Context, in *catalog.CreateProductRequest) (*catalog.CreateProductResponse, error) {
	l := logic.NewCreateProductLogic(ctx, s.svcCtx, s.universe)
	return l.CreateProduct(in)
}

func (s *CatalogClientServer) UpdateProduct(ctx context.Context, in *catalog.UpdateProductRequest) (*catalog.UpdateProductResponse, error) {
	l := logic.NewUpdateProductLogic(ctx, s.svcCtx, s.universe)
	return l.UpdateProduct(in)
}

func (s *CatalogClientServer) DeleteProduct(ctx context.Context, in *catalog.DeleteProductRequest) (*catalog.DeleteProductResponse, error) {
	l := logic.NewDeleteProductLogic(ctx, s.svcCtx, s.universe)
	return l.DeleteProduct(in)
}