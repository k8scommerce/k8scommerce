package products

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/helpers"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/types"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/catalogclient"

	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllProductsLogic {
	return GetAllProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllProductsLogic) GetAllProducts(req types.GetAllProductsRequest) (resp *types.GetAllProductsResponse, err error) {
	response, err := l.svcCtx.CatalogRpc.GetAllProducts(l.ctx, &catalogclient.GetAllProductsRequest{
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
		Filter:      l.ctx.Value(types.Filter).(string),
		StoreId:     l.ctx.Value(types.StoreKey).(int64),
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	resp = &types.GetAllProductsResponse{}
	utils.TransformObj(response, &resp)

	// format the currency to the locale and language
	for i := 0; i < len(resp.Products); i++ {
		for x := 0; x < len(resp.Products[i].Variants); x++ {
			if resp.Products[i].Variants[x].Price != (types.Price{}) {
				helpers.ConvertOutgoingPrices(l.ctx, &resp.Products[i].Variants[x].Price)
			}
		}
	}
	return resp, err
}
