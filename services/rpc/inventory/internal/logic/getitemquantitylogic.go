package logic

import (
	"context"
	"sync"

	"k8scommerce/services/rpc/inventory/internal/svc"
	"k8scommerce/services/rpc/inventory/pb/inventory"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
)

type galaxyGetItemQuantityLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetItemQuantityLogic *galaxyGetItemQuantityLogicHelper

type GetItemQuantityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetItemQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetItemQuantityLogic {
	return &GetItemQuantityLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetItemQuantityLogic) GetItemQuantity(in *inventory.GetItemQuantityRequest) (*inventory.GetItemQuantityResponse, error) {

	res := &inventory.GetItemQuantityResponse{
		Sku:     in.Sku,
		Quanity: 100,
	}

	return res, nil

}
