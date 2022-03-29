// Code generated by goctl. DO NOT EDIT!
// Source: cart.proto

package cartclient

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddItemRequest              = cart.AddItemRequest
	AttachCustomerRequest       = cart.AttachCustomerRequest
	BulkAddItemsRequest         = cart.BulkAddItemsRequest
	Cart                        = cart.Cart
	CartResponse                = cart.CartResponse
	ClearCartRequest            = cart.ClearCartRequest
	CreateCartRequest           = cart.CreateCartRequest
	GetByCartIdRequest          = cart.GetByCartIdRequest
	GetBySessionIdRequest       = cart.GetBySessionIdRequest
	Item                        = cart.Item
	OthersBought                = cart.OthersBought
	RemoveItemRequest           = cart.RemoveItemRequest
	SimilarProducts             = cart.SimilarProducts
	UpdateCustomerDetailRequest = cart.UpdateCustomerDetailRequest
	UpdateItemQuantityRequest   = cart.UpdateItemQuantityRequest
	UpdateStatusRequest         = cart.UpdateStatusRequest

	CartClient interface {
		CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CartResponse, error)
		AttachCustomer(ctx context.Context, in *AttachCustomerRequest, opts ...grpc.CallOption) (*CartResponse, error)
		UpdateCustomerDetail(ctx context.Context, in *UpdateCustomerDetailRequest, opts ...grpc.CallOption) (*CartResponse, error)
		UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*CartResponse, error)
		GetByCartId(ctx context.Context, in *GetByCartIdRequest, opts ...grpc.CallOption) (*CartResponse, error)
		GetBySessionId(ctx context.Context, in *GetBySessionIdRequest, opts ...grpc.CallOption) (*CartResponse, error)
		AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*CartResponse, error)
		BulkAddItems(ctx context.Context, in *BulkAddItemsRequest, opts ...grpc.CallOption) (*CartResponse, error)
		UpdateItemQuantity(ctx context.Context, in *UpdateItemQuantityRequest, opts ...grpc.CallOption) (*CartResponse, error)
		RemoveItem(ctx context.Context, in *RemoveItemRequest, opts ...grpc.CallOption) (*CartResponse, error)
		ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*CartResponse, error)
	}

	defaultCartClient struct {
		cli zrpc.Client
	}
)

func NewCartClient(cli zrpc.Client) CartClient {
	return &defaultCartClient{
		cli: cli,
	}
}

func (m *defaultCartClient) CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.CreateCart(ctx, in, opts...)
}

func (m *defaultCartClient) AttachCustomer(ctx context.Context, in *AttachCustomerRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.AttachCustomer(ctx, in, opts...)
}

func (m *defaultCartClient) UpdateCustomerDetail(ctx context.Context, in *UpdateCustomerDetailRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.UpdateCustomerDetail(ctx, in, opts...)
}

func (m *defaultCartClient) UpdateStatus(ctx context.Context, in *UpdateStatusRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.UpdateStatus(ctx, in, opts...)
}

func (m *defaultCartClient) GetByCartId(ctx context.Context, in *GetByCartIdRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.GetByCartId(ctx, in, opts...)
}

func (m *defaultCartClient) GetBySessionId(ctx context.Context, in *GetBySessionIdRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.GetBySessionId(ctx, in, opts...)
}

func (m *defaultCartClient) AddItem(ctx context.Context, in *AddItemRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.AddItem(ctx, in, opts...)
}

func (m *defaultCartClient) BulkAddItems(ctx context.Context, in *BulkAddItemsRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.BulkAddItems(ctx, in, opts...)
}

func (m *defaultCartClient) UpdateItemQuantity(ctx context.Context, in *UpdateItemQuantityRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.UpdateItemQuantity(ctx, in, opts...)
}

func (m *defaultCartClient) RemoveItem(ctx context.Context, in *RemoveItemRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.RemoveItem(ctx, in, opts...)
}

func (m *defaultCartClient) ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.ClearCart(ctx, in, opts...)
}
