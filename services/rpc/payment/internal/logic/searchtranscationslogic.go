package logic

import (
	"context"

	"k8scommerce/services/rpc/payment/internal/svc"
	"k8scommerce/services/rpc/payment/pb/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTranscationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTranscationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTranscationsLogic {
	return &SearchTranscationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTranscationsLogic) SearchTranscations(in *payment.SearchTransactionsRequest) (*payment.SearchTransactionsResponse, error) {
	res := &payment.SearchTransactionsResponse{}
	return res, nil
}
