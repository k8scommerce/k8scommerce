package logic

import (
	"context"
	"sync"

	"github.com/k8s-commerce/k8s-commerce/services/rpc/payment/internal/svc"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/payment/pb/payment"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxySearchTranscationsLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entrySearchTranscationsLogic *galaxySearchTranscationsLogicHelper

type SearchTranscationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewSearchTranscationsLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *SearchTranscationsLogic {
	return &SearchTranscationsLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *SearchTranscationsLogic) SearchTranscations(in *payment.SearchTransactionsRequest) (*payment.SearchTransactionsResponse, error) {
	res := &payment.SearchTransactionsResponse{}
	return res, nil
}
