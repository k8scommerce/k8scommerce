package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/payment/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/payment/pb/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTranscationByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTranscationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTranscationByIdLogic {
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
