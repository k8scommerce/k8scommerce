// Code generated by goctl. DO NOT EDIT!
// Source: inventory.proto

package inventoryclient

import (
	"context"

	"k8scommerce/services/rpc/inventory/pb/inventory"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Brand                   = inventory.Brand
	GetItemQuantityRequest  = inventory.GetItemQuantityRequest
	GetItemQuantityResponse = inventory.GetItemQuantityResponse
	InventoryItem           = inventory.InventoryItem
	SockLevel               = inventory.SockLevel
	Supplier                = inventory.Supplier
	Warehouse               = inventory.Warehouse

	InventoryClient interface {
		GetItemQuantity(ctx context.Context, in *GetItemQuantityRequest, opts ...grpc.CallOption) (*GetItemQuantityResponse, error)
	}

	defaultInventoryClient struct {
		cli zrpc.Client
	}
)

func NewInventoryClient(cli zrpc.Client) InventoryClient {
	return &defaultInventoryClient{
		cli: cli,
	}
}

func (m *defaultInventoryClient) GetItemQuantity(ctx context.Context, in *GetItemQuantityRequest, opts ...grpc.CallOption) (*GetItemQuantityResponse, error) {
	client := inventory.NewInventoryClientClient(m.cli.Conn())
	return client.GetItemQuantity(ctx, in, opts...)
}
