package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/warehouse/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/warehouse/pb/warehouse"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWarehouseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateWarehouseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWarehouseLogic {
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
