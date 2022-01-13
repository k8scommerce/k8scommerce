package logic

import (
	"context"
	"encoding/json"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
	"k8scommerce/services/rpc/catalog/catalogclient"

	"github.com/tal-tech/go-zero/core/logx"
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
		CategoryId:  req.CategoryId,
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
		SortOn:      req.SortOn,
		StoreId:     l.ctx.Value(types.StoreKey).(int64),
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	res := &types.GetProductsByCategoryIdResponse{}
	b, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, res)

	// format the currency to the locale and language
	for i := 0; i < len(res.Products); i++ {
		for x := 0; x < len(res.Products[i].Variants); x++ {
			if res.Products[i].Variants[x].Price != (types.Price{}) {
				convertOutgoingPrices(l.ctx, &res.Products[i].Variants[x].Price)
			}
		}
	}

	return res, err
}
