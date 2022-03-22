package products

import (
	"context"

	"k8scommerce/internal/utils"
	"k8scommerce/services/api/client/helpers"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
	"k8scommerce/services/rpc/catalog/catalogclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductsByCategoryIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductsByCategoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductsByCategoryIdLogic {
	return GetProductsByCategoryIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductsByCategoryIdLogic) GetProductsByCategoryId(req types.GetProductsByCategoryIdRequest) (resp *types.GetProductsByCategoryIdResponse, err error) {
	response, err := l.svcCtx.CatalogRpc.GetProductsByCategoryId(l.ctx, &catalogclient.GetProductsByCategoryIdRequest{
		StoreId:     l.ctx.Value(types.StoreKey).(int64),
		CategoryId:  req.CategoryId,
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
		Filter:      l.ctx.Value(types.Filter).(string),
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	resp = &types.GetProductsByCategoryIdResponse{}
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
