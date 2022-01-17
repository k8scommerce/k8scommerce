package logic

import (
	"context"
	"encoding/json"
	"net/http"

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

func (l *GetProductBySlugLogic) GetProductBySlug(req types.GetProductBySlugRequest) (resp *types.GetProductResponse, err error) {
	resp = &types.GetProductResponse{}
	l.Logger.Info(req)
	getProductBySlugResponse, err := l.svcCtx.CatalogRpc.GetProductBySlug(l.ctx, &catalogclient.GetProductBySlugRequest{
		Slug:    req.Slug,
		StoreId: l.ctx.Value(types.StoreKey).(int64),
	})
	if err != nil {
		resp.ResponseStatus = httpResponse(http.StatusBadRequest, err.Error())
		return resp, err
	}

	l.Logger.Info(getProductBySlugResponse)

	// convert from one type to another
	// the structs are identical
	resp.Product = types.Product{}
	b, err := json.Marshal(getProductBySlugResponse.Product)
	if err != nil {
		resp.ResponseStatus = httpResponse(http.StatusBadRequest, err.Error())
		return resp, err
	}
	err = json.Unmarshal(b, &resp.Product)

	// format the currency to the locale and language
	for x := 0; x < len(resp.Product.Variants); x++ {
		if resp.Product.Variants[x].Price != (types.Price{}) {
			convertOutgoingPrices(l.ctx, &resp.Product.Variants[x].Price)
		}
	}

	// looks good
	resp.ResponseStatus = httpResponse(http.StatusOK)
	return resp, err
}
