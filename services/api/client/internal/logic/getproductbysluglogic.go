package logic

import (
	"context"
	"encoding/json"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
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
	getProductBySlugResponse, err := l.svcCtx.CatalogRpc.GetProductBySlug(l.ctx, &catalogclient.GetProductBySlugRequest{
		Slug:    req.Slug,
		StoreId: l.ctx.Value(types.StoreKey).(int64),
	})
	if err != nil {
		return nil, err
	}

	l.Logger.Info(getProductBySlugResponse)

	// convert from one type to another
	// the structs are identical
	res := &types.Product{}
	b, err := json.Marshal(getProductBySlugResponse.Product)
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
