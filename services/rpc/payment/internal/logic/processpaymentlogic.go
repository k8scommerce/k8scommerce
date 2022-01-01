package logic

import (
	"context"
	"sync"

	"github.com/k8s-commerce/k8s-commerce/services/rpc/payment/internal/svc"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/payment/pb/payment"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyProcessPaymentLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryProcessPaymentLogic *galaxyProcessPaymentLogicHelper

type ProcessPaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewProcessPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *ProcessPaymentLogic {
	return &ProcessPaymentLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *ProcessPaymentLogic) ProcessPayment(in *payment.ProcessPaymentRequest) (*payment.ProcessPaymentResponse, error) {
	res := &payment.ProcessPaymentResponse{}
	return res, nil
}
