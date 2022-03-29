package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/inventory/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/inventory/pb/inventory"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetItemsQuantityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetItemsQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetItemsQuantityLogic {
	return &GetItemsQuantityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetItemsQuantityLogic) GetItemsQuantity(in *inventory.GetItemsQuantityRequest) (*inventory.GetItemsQuantityResponse, error) {
	res := &inventory.GetItemsQuantityResponse{
		StockLevels: []*inventory.StockLevel{},
	}
	for _, sku := range in.Skus {
		res.StockLevels = append(res.StockLevels, &inventory.StockLevel{
			StoreId:     in.StoreId,
			WarehouseId: 1,
			Sku:         sku,
			Row:         "10",
			Shelf:       "2",
			Bin:         "12b",
			Quantity:    100,
		})
	}
	return res, nil
}
