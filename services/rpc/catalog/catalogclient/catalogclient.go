// Code generated by goctl. DO NOT EDIT!
// Source: catalog.proto

package catalogclient

import (
	"context"

	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Asset                             = catalog.Asset
	Category                          = catalog.Category
	CreateCategoryRequest             = catalog.CreateCategoryRequest
	CreateCategoryResponse            = catalog.CreateCategoryResponse
	CreateProductRequest              = catalog.CreateProductRequest
	CreateProductResponse             = catalog.CreateProductResponse
	DeleteCategoryRequest             = catalog.DeleteCategoryRequest
	DeleteCategoryResponse            = catalog.DeleteCategoryResponse
	DeleteProductRequest              = catalog.DeleteProductRequest
	DeleteProductResponse             = catalog.DeleteProductResponse
	GetAllCategoriesRequest           = catalog.GetAllCategoriesRequest
	GetAllCategoriesResponse          = catalog.GetAllCategoriesResponse
	GetAllProductsRequest             = catalog.GetAllProductsRequest
	GetAllProductsResponse            = catalog.GetAllProductsResponse
	GetCategoryByIdRequest            = catalog.GetCategoryByIdRequest
	GetCategoryByIdResponse           = catalog.GetCategoryByIdResponse
	GetCategoryBySlugRequest          = catalog.GetCategoryBySlugRequest
	GetCategoryBySlugResponse         = catalog.GetCategoryBySlugResponse
	GetProductByIdRequest             = catalog.GetProductByIdRequest
	GetProductByIdResponse            = catalog.GetProductByIdResponse
	GetProductBySkuRequest            = catalog.GetProductBySkuRequest
	GetProductBySkuResponse           = catalog.GetProductBySkuResponse
	GetProductBySlugRequest           = catalog.GetProductBySlugRequest
	GetProductBySlugResponse          = catalog.GetProductBySlugResponse
	GetProductsByCategoryIdRequest    = catalog.GetProductsByCategoryIdRequest
	GetProductsByCategoryIdResponse   = catalog.GetProductsByCategoryIdResponse
	GetProductsByCategorySlugRequest  = catalog.GetProductsByCategorySlugRequest
	GetProductsByCategorySlugResponse = catalog.GetProductsByCategorySlugResponse
	Price                             = catalog.Price
	Product                           = catalog.Product
	UpdateCategoryRequest             = catalog.UpdateCategoryRequest
	UpdateCategoryResponse            = catalog.UpdateCategoryResponse
	UpdateProductRequest              = catalog.UpdateProductRequest
	UpdateProductResponse             = catalog.UpdateProductResponse
	UploadAssetRequest                = catalog.UploadAssetRequest
	Variant                           = catalog.Variant

	CatalogClient interface {
		//  categories
		GetAllCategories(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error)
		GetCategoryBySlug(ctx context.Context, in *GetCategoryBySlugRequest, opts ...grpc.CallOption) (*GetCategoryBySlugResponse, error)
		GetCategoryById(ctx context.Context, in *GetCategoryByIdRequest, opts ...grpc.CallOption) (*GetCategoryByIdResponse, error)
		CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
		UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error)
		DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error)
		//  products
		GetProductBySku(ctx context.Context, in *GetProductBySkuRequest, opts ...grpc.CallOption) (*GetProductBySkuResponse, error)
		GetProductBySlug(ctx context.Context, in *GetProductBySlugRequest, opts ...grpc.CallOption) (*GetProductBySlugResponse, error)
		GetProductById(ctx context.Context, in *GetProductByIdRequest, opts ...grpc.CallOption) (*GetProductByIdResponse, error)
		GetProductsByCategoryId(ctx context.Context, in *GetProductsByCategoryIdRequest, opts ...grpc.CallOption) (*GetProductsByCategoryIdResponse, error)
		GetProductsByCategorySlug(ctx context.Context, in *GetProductsByCategorySlugRequest, opts ...grpc.CallOption) (*GetProductsByCategorySlugResponse, error)
		GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error)
		CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
		UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error)
		DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
		UploadAsset(ctx context.Context, opts ...grpc.CallOption) (catalog.CatalogClient_UploadAssetClient, error)
	}

	defaultCatalogClient struct {
		cli zrpc.Client
	}
)

func NewCatalogClient(cli zrpc.Client) CatalogClient {
	return &defaultCatalogClient{
		cli: cli,
	}
}

//  categories
func (m *defaultCatalogClient) GetAllCategories(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.GetAllCategories(ctx, in, opts...)
}

func (m *defaultCatalogClient) GetCategoryBySlug(ctx context.Context, in *GetCategoryBySlugRequest, opts ...grpc.CallOption) (*GetCategoryBySlugResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.GetCategoryBySlug(ctx, in, opts...)
}

func (m *defaultCatalogClient) GetCategoryById(ctx context.Context, in *GetCategoryByIdRequest, opts ...grpc.CallOption) (*GetCategoryByIdResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.GetCategoryById(ctx, in, opts...)
}

func (m *defaultCatalogClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.CreateCategory(ctx, in, opts...)
}

func (m *defaultCatalogClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.UpdateCategory(ctx, in, opts...)
}

func (m *defaultCatalogClient) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.DeleteCategory(ctx, in, opts...)
}

//  products
func (m *defaultCatalogClient) GetProductBySku(ctx context.Context, in *GetProductBySkuRequest, opts ...grpc.CallOption) (*GetProductBySkuResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.GetProductBySku(ctx, in, opts...)
}

func (m *defaultCatalogClient) GetProductBySlug(ctx context.Context, in *GetProductBySlugRequest, opts ...grpc.CallOption) (*GetProductBySlugResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.GetProductBySlug(ctx, in, opts...)
}

func (m *defaultCatalogClient) GetProductById(ctx context.Context, in *GetProductByIdRequest, opts ...grpc.CallOption) (*GetProductByIdResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.GetProductById(ctx, in, opts...)
}

func (m *defaultCatalogClient) GetProductsByCategoryId(ctx context.Context, in *GetProductsByCategoryIdRequest, opts ...grpc.CallOption) (*GetProductsByCategoryIdResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.GetProductsByCategoryId(ctx, in, opts...)
}

func (m *defaultCatalogClient) GetProductsByCategorySlug(ctx context.Context, in *GetProductsByCategorySlugRequest, opts ...grpc.CallOption) (*GetProductsByCategorySlugResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.GetProductsByCategorySlug(ctx, in, opts...)
}

func (m *defaultCatalogClient) GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.GetAllProducts(ctx, in, opts...)
}

func (m *defaultCatalogClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.CreateProduct(ctx, in, opts...)
}

func (m *defaultCatalogClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.UpdateProduct(ctx, in, opts...)
}

func (m *defaultCatalogClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.DeleteProduct(ctx, in, opts...)
}

func (m *defaultCatalogClient) UploadAsset(ctx context.Context, opts ...grpc.CallOption) (catalog.CatalogClient_UploadAssetClient, error) {
	client := catalog.NewCatalogClientClient(m.cli.Conn())
	return client.UploadAsset(ctx, opts...)
}
