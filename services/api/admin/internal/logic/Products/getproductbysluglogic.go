package products

import (
	"context"
	"encoding/json"

	"k8scommerce/services/api/admin/internal/helpers"
	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
	"k8scommerce/services/rpc/catalog/catalogclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductBySlugLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductBySlugLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductBySlugLogic {
	return GetProductBySlugLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductBySlugLogic) GetProductBySlug(req types.GetProductBySlugRequest) (resp *types.Product, err error) {
	resp = &types.Product{}

	response, err := l.svcCtx.CatalogRpc.GetProductBySlug(l.ctx, &catalogclient.GetProductBySlugRequest{
		Slug:    req.Slug,
		StoreId: l.ctx.Value(types.StoreKey).(int64),
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	b, err := json.Marshal(response.Product)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, resp)

	// format the currency to the locale and language
	for x := 0; x < len(resp.Variants); x++ {
		if resp.Variants[x].Price != (types.Price{}) {
			helpers.ConvertOutgoingPrices(l.ctx, &resp.Variants[x].Price)
		}
	}

	return resp, err
}