package logic

import (
	"context"
	"encoding/json"

	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
	"k8scommerce/services/rpc/catalog/catalogclient"

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
	getOneBySkuResponse, err := l.svcCtx.CatalogRpc.GetProductBySku(l.ctx, &catalogclient.GetProductBySkuRequest{
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

	// format the currency to the locale and language
	for x := 0; x < len(res.Variants); x++ {
		if res.Variants[x].Price != (types.Price{}) {
			convertOutgoingPrices(l.ctx, &res.Variants[x].Price)
		}
	}
	return res, err
}
