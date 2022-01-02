package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/client/pb/shipping"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type GetQuoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetQuoteLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetQuoteLogic {
	return &GetQuoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetQuoteLogic) GetQuote(in *shipping.GetQuoteRequest) (*shipping.GetQuoteResponse, error) {
	res := &shipping.GetQuoteResponse{}
	return res, nil
}
