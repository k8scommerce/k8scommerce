package logic

import (
	"context"
	"encoding/json"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
	"k8scommerce/services/rpc/catalog/catalogclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllProductsRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllProductsRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllProductsRequestLogic {
	return GetAllProductsRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllProductsRequestLogic) GetAllProductsRequest(req types.GetAllProductsRequest) (*types.GetAllProductsResponse, error) {
	response, err := l.svcCtx.CatalogRpc.GetAllProducts(l.ctx, &catalogclient.GetAllProductsRequest{
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
		SortOn:      req.SortOn,
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
				convertOutgoingPrices(l.ctx, &res.Products[i].Variants[x].Price)
			}
		}
	}
	return res, err
}
