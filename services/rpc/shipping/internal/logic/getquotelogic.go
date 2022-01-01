package logic

import (
	"context"
	"shipping/internal/svc"
	"shipping/shipping"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetQuoteLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetQuoteLogic *galaxyGetQuoteLogicHelper

type GetQuoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetQuoteLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetQuoteLogic {
	return &GetQuoteLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetQuoteLogic) GetQuote(in *shipping.GetQuoteRequest) (*shipping.GetQuoteResponse, error) {
	res := &shipping.GetQuoteResponse{}
	return res, nil
}
