// Code generated by goctl. DO NOT EDIT!
// Source: store.proto

package storeclient

import (
	"context"

	"k8scommerce/services/rpc/store/pb/store"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateStoreRequest   = store.CreateStoreRequest
	CreateStoreResponse  = store.CreateStoreResponse
	GetAllStoresRequest  = store.GetAllStoresRequest
	GetAllStoresResponse = store.GetAllStoresResponse
	GetStoreByIdRequest  = store.GetStoreByIdRequest
	GetStoreByIdResponse = store.GetStoreByIdResponse
	Store                = store.Store
	StoreAddress         = store.StoreAddress
	StoreSetting         = store.StoreSetting

	StoreClient interface {
		CreateStore(ctx context.Context, in *CreateStoreRequest, opts ...grpc.CallOption) (*CreateStoreResponse, error)
		GetStoreById(ctx context.Context, in *GetStoreByIdRequest, opts ...grpc.CallOption) (*GetStoreByIdResponse, error)
		GetAllStores(ctx context.Context, in *GetAllStoresRequest, opts ...grpc.CallOption) (*GetAllStoresResponse, error)
	}

	defaultStoreClient struct {
		cli zrpc.Client
	}
)

func NewStoreClient(cli zrpc.Client) StoreClient {
	return &defaultStoreClient{
		cli: cli,
	}
}

func (m *defaultStoreClient) CreateStore(ctx context.Context, in *CreateStoreRequest, opts ...grpc.CallOption) (*CreateStoreResponse, error) {
	client := store.NewStoreClientClient(m.cli.Conn())
	return client.CreateStore(ctx, in, opts...)
}

func (m *defaultStoreClient) GetStoreById(ctx context.Context, in *GetStoreByIdRequest, opts ...grpc.CallOption) (*GetStoreByIdResponse, error) {
	client := store.NewStoreClientClient(m.cli.Conn())
	return client.GetStoreById(ctx, in, opts...)
}

func (m *defaultStoreClient) GetAllStores(ctx context.Context, in *GetAllStoresRequest, opts ...grpc.CallOption) (*GetAllStoresResponse, error) {
	client := store.NewStoreClientClient(m.cli.Conn())
	return client.GetAllStores(ctx, in, opts...)
}
