package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/client/pb/payment"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type SearchTranscationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTranscationsLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *SearchTranscationsLogic {
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
