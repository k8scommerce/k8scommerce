package logic

import (
	"context"

	"k8scommerce/services/rpc/warehouse/internal/svc"
	"k8scommerce/services/rpc/warehouse/pb/warehouse"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWarehouseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateWarehouseLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *CreateWarehouseLogic {
	return &CreateWarehouseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateWarehouseLogic) CreateWarehouse(in *warehouse.CreateWarehouseRequest) (*warehouse.CreateWarehouseResponse, error) {
	res := &warehouse.CreateWarehouseResponse{}
	return res, nil
}
