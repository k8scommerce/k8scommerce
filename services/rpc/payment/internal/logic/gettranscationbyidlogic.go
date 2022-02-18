package logic

import (
	"context"

	"k8scommerce/services/rpc/payment/internal/svc"
	"k8scommerce/services/rpc/payment/pb/payment"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTranscationByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTranscationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetTranscationByIdLogic {
	return &GetTranscationByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTranscationByIdLogic) GetTranscationById(in *payment.GetTranscationByIdRequest) (*payment.GetTranscationByIdResponse, error) {
	res := &payment.GetTranscationByIdResponse{}
	return res, nil

}
