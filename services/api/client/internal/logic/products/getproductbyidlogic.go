package products

import (
	"context"
	"encoding/json"

	"k8scommerce/services/api/client/helpers"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
	"k8scommerce/services/rpc/catalog/catalogclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductByIdLogic {
	return GetProductByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductByIdLogic) GetProductById(req types.GetProductByIdRequest) (resp *types.GetProductResponse, err error) {
	resp = &types.GetProductResponse{}

	getOneByIdResponse, err := l.svcCtx.CatalogRpc.GetProductById(l.ctx, &catalogclient.GetProductByIdRequest{
		Id:      req.Id,
		StoreId: l.ctx.Value(types.StoreKey).(int64),
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	prod := &types.Product{}
	b, err := json.Marshal(getOneByIdResponse.Product)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, prod)

	// format the currency to the locale and language
	for x := 0; x < len(prod.Variants); x++ {
		if prod.Variants[x].Price != (types.Price{}) {
			helpers.ConvertOutgoingPrices(l.ctx, &prod.Variants[x].Price)
		}
	}

	resp.Product = *prod
	return resp, err
}
