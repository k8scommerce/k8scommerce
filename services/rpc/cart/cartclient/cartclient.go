// Code generated by goctl. DO NOT EDIT!
// Source: cart.proto

package cartclient

import (
	"context"

	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddItemToCartRequest             = cart.AddItemToCartRequest
	AddItemToCartResponse            = cart.AddItemToCartResponse
	Cart                             = cart.Cart
	ClearCartRequest                 = cart.ClearCartRequest
	ClearCartResponse                = cart.ClearCartResponse
	GetCartRequest                   = cart.GetCartRequest
	GetCartResponse                  = cart.GetCartResponse
	Item                             = cart.Item
	OthersBought                     = cart.OthersBought
	Price                            = cart.Price
	RemoveItemInCartRequest          = cart.RemoveItemInCartRequest
	RemoveItemInCartResponse         = cart.RemoveItemInCartResponse
	SimilarProducts                  = cart.SimilarProducts
	UpdateItemQuantityInCartRequest  = cart.UpdateItemQuantityInCartRequest
	UpdateItemQuantityInCartResponse = cart.UpdateItemQuantityInCartResponse

	CartClient interface {
		GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error)
		ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartResponse, error)
		AddItemToCart(ctx context.Context, in *AddItemToCartRequest, opts ...grpc.CallOption) (*AddItemToCartResponse, error)
		UpdateItemQuantityInCart(ctx context.Context, in *UpdateItemQuantityInCartRequest, opts ...grpc.CallOption) (*UpdateItemQuantityInCartResponse, error)
		RemoveItemInCart(ctx context.Context, in *RemoveItemInCartRequest, opts ...grpc.CallOption) (*RemoveItemInCartResponse, error)
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

func (m *defaultCartClient) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.GetCart(ctx, in, opts...)
}

func (m *defaultCartClient) ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.ClearCart(ctx, in, opts...)
}

func (m *defaultCartClient) AddItemToCart(ctx context.Context, in *AddItemToCartRequest, opts ...grpc.CallOption) (*AddItemToCartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.AddItemToCart(ctx, in, opts...)
}

func (m *defaultCartClient) UpdateItemQuantityInCart(ctx context.Context, in *UpdateItemQuantityInCartRequest, opts ...grpc.CallOption) (*UpdateItemQuantityInCartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.UpdateItemQuantityInCart(ctx, in, opts...)
}

func (m *defaultCartClient) RemoveItemInCart(ctx context.Context, in *RemoveItemInCartRequest, opts ...grpc.CallOption) (*RemoveItemInCartResponse, error) {
	client := cart.NewCartClientClient(m.cli.Conn())
	return client.RemoveItemInCart(ctx, in, opts...)
}
