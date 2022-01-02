package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/client/pb/payment"
	"github.com/localrivet/galaxycache"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetTransactionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTransactionsLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetTransactionsLogic {
	return &GetTransactionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTransactionsLogic) GetTransactions(in *payment.ProcessPaymentRequest) (*payment.ProcessPaymentResponse, error) {
	res := &payment.ProcessPaymentResponse{}
	return res, nil

}
