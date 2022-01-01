// Code generated by goctl. DO NOT EDIT!
// Source: similarproducts.proto

package similarproductsclient

import (
	"context"

	"github.com/k8s-commerce/k8s-commerce/services/rpc/similarproducts/pb/similarproducts"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetSimilarProductsBySkuRequest  = similarproducts.GetSimilarProductsBySkuRequest
	GetSimilarProductsBySkuResponse = similarproducts.GetSimilarProductsBySkuResponse

	SimilarProductsClient interface {
		GetSimilarProductsBySku(ctx context.Context, in *GetSimilarProductsBySkuRequest, opts ...grpc.CallOption) (*GetSimilarProductsBySkuResponse, error)
	}

	defaultSimilarProductsClient struct {
		cli zrpc.Client
	}
)

func NewSimilarProductsClient(cli zrpc.Client) SimilarProductsClient {
	return &defaultSimilarProductsClient{
		cli: cli,
	}
}

func (m *defaultSimilarProductsClient) GetSimilarProductsBySku(ctx context.Context, in *GetSimilarProductsBySkuRequest, opts ...grpc.CallOption) (*GetSimilarProductsBySkuResponse, error) {
	client := similarproducts.NewSimilarProductsClientClient(m.cli.Conn())
	return client.GetSimilarProductsBySku(ctx, in, opts...)
}
