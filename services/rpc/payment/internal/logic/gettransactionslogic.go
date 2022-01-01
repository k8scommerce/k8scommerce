package logic

import (
	"context"
	"sync"

	"github.com/k8s-commerce/k8s-commerce/services/rpc/payment/internal/svc"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/payment/pb/payment"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetTransactionsLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetTransactionsLogic *galaxyGetTransactionsLogicHelper

type GetTransactionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetTransactionsLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetTransactionsLogic {
	return &GetTransactionsLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetTransactionsLogic) GetTransactions(in *payment.ProcessPaymentRequest) (*payment.ProcessPaymentResponse, error) {
	res := &payment.ProcessPaymentResponse{}
	return res, nil
}
