package logic

import (
	"context"
	"encoding/json"

	"ecomm/services/api/admin/internal/svc"
	"ecomm/services/api/admin/internal/types"
	"ecomm/services/rpc/product/productclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductBySlugRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductBySlugRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductBySlugRequestLogic {
	return GetProductBySlugRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductBySlugRequestLogic) GetProductBySlugRequest(req types.GetProductBySlugRequest) (*types.Product, error) {
	getOneBySlugResponse, err := l.svcCtx.ProductRpc.GetProductBySlug(l.ctx, &productclient.GetProductBySlugRequest{
		Slug: req.Slug,
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	res := &types.Product{}
	b, err := json.Marshal(getOneBySlugResponse.Product)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, res)
	return res, err
}
