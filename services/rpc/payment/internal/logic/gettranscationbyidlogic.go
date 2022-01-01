package logic

import (
	"context"
	"sync"

	"github.com/k8s-commerce/k8s-commerce/services/rpc/payment/internal/svc"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/payment/pb/payment"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetTranscationByIdLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetTranscationByIdLogic *galaxyGetTranscationByIdLogicHelper

type GetTranscationByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetTranscationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetTranscationByIdLogic {
	return &GetTranscationByIdLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetTranscationByIdLogic) GetTranscationById(in *payment.GetTranscationByIdRequest) (*payment.GetTranscationByIdResponse, error) {
	res := &payment.GetTranscationByIdResponse{}
	return res, nil
}
