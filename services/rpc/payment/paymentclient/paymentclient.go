// Code generated by goctl. DO NOT EDIT!
// Source: payment.proto

package paymentclient

import (
	"context"

	"k8scommerce/services/rpc/payment/pb/payment"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreditCard                 = payment.CreditCard
	GetTransactionsRequest     = payment.GetTransactionsRequest
	GetTransactionsResponse    = payment.GetTransactionsResponse
	GetTranscationByIdRequest  = payment.GetTranscationByIdRequest
	GetTranscationByIdResponse = payment.GetTranscationByIdResponse
	ProcessPaymentRequest      = payment.ProcessPaymentRequest
	ProcessPaymentResponse     = payment.ProcessPaymentResponse
	SearchTransactionsRequest  = payment.SearchTransactionsRequest
	SearchTransactionsResponse = payment.SearchTransactionsResponse
	Transaction                = payment.Transaction

	PaymentClient interface {
		ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error)
		GetTransactions(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error)
		GetTranscationById(ctx context.Context, in *GetTranscationByIdRequest, opts ...grpc.CallOption) (*GetTranscationByIdResponse, error)
		SearchTranscations(ctx context.Context, in *SearchTransactionsRequest, opts ...grpc.CallOption) (*SearchTransactionsResponse, error)
	}

	defaultPaymentClient struct {
		cli zrpc.Client
	}
)

func NewPaymentClient(cli zrpc.Client) PaymentClient {
	return &defaultPaymentClient{
		cli: cli,
	}
}

func (m *defaultPaymentClient) ProcessPayment(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error) {
	client := payment.NewPaymentClientClient(m.cli.Conn())
	return client.ProcessPayment(ctx, in, opts...)
}

func (m *defaultPaymentClient) GetTransactions(ctx context.Context, in *ProcessPaymentRequest, opts ...grpc.CallOption) (*ProcessPaymentResponse, error) {
	client := payment.NewPaymentClientClient(m.cli.Conn())
	return client.GetTransactions(ctx, in, opts...)
}

func (m *defaultPaymentClient) GetTranscationById(ctx context.Context, in *GetTranscationByIdRequest, opts ...grpc.CallOption) (*GetTranscationByIdResponse, error) {
	client := payment.NewPaymentClientClient(m.cli.Conn())
	return client.GetTranscationById(ctx, in, opts...)
}

func (m *defaultPaymentClient) SearchTranscations(ctx context.Context, in *SearchTransactionsRequest, opts ...grpc.CallOption) (*SearchTransactionsResponse, error) {
	client := payment.NewPaymentClientClient(m.cli.Conn())
	return client.SearchTranscations(ctx, in, opts...)
}
