package logic

import (
	"context"
	"encoding/json"

	"ecomm/services/api/admin/internal/svc"
	"ecomm/services/api/admin/internal/types"
	"ecomm/services/rpc/product/productclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductBySkuRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductBySkuRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductBySkuRequestLogic {
	return GetProductBySkuRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductBySkuRequestLogic) GetProductBySkuRequest(req types.GetProductBySkuRequest) (*types.Product, error) {
	getOneBySkuResponse, err := l.svcCtx.ProductRpc.GetProductBySku(l.ctx, &productclient.GetProductBySkuRequest{
		Sku: req.Sku,
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	res := &types.Product{}
	b, err := json.Marshal(getOneBySkuResponse.Product)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, res)
	return res, err
}
