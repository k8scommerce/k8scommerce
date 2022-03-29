package logic

import (
	"context"
	"sync"

	"github.com/k8scommerce/k8scommerce/services/rpc/inventory/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/inventory/pb/inventory"

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
}

func NewGetItemQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetItemQuantityLogic {
	return &GetItemQuantityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetItemQuantityLogic) GetItemQuantity(in *inventory.GetItemQuantityRequest) (*inventory.GetItemQuantityResponse, error) {
	res := &inventory.GetItemQuantityResponse{
		StockLevel: &inventory.StockLevel{
			StoreId:     in.StoreId,
			WarehouseId: 1,
			Sku:         in.Sku,
			Row:         "10",
			Shelf:       "2",
			Bin:         "12b",
			Quantity:    100,
		},
	}

	return res, nil
}
