// Code generated by goctl. DO NOT EDIT!
// Source: shipping.proto

package shippingclient

import (
	"context"

	"k8scommerce/services/rpc/shipping/pb/shipping"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	From             = shipping.From
	GetQuoteRequest  = shipping.GetQuoteRequest
	GetQuoteResponse = shipping.GetQuoteResponse
	To               = shipping.To

	ShippingClient interface {
		GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*GetQuoteResponse, error)
	}

	defaultShippingClient struct {
		cli zrpc.Client
	}
)

func NewShippingClient(cli zrpc.Client) ShippingClient {
	return &defaultShippingClient{
		cli: cli,
	}
}

func (m *defaultShippingClient) GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*GetQuoteResponse, error) {
	client := shipping.NewShippingClientClient(m.cli.Conn())
	return client.GetQuote(ctx, in, opts...)
}
