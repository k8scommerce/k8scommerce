package products

import (
	"context"
	"encoding/json"

	"k8scommerce/services/api/client/helpers"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
	"k8scommerce/services/rpc/catalog/catalogclient"

	"github.com/tal-tech/go-zero/core/logx"
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
	res := &types.GetAllProductsResponse{}
	b, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &res)

	// format the currency to the locale and language
	for i := 0; i < len(res.Products); i++ {
		for x := 0; x < len(res.Products[i].Variants); x++ {
			if res.Products[i].Variants[x].Price != (types.Price{}) {
				helpers.ConvertOutgoingPrices(l.ctx, &res.Products[i].Variants[x].Price)
			}
		}
	}
	return res, err
}
