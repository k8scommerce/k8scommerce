package products

import (
	"context"

	"k8scommerce/internal/utils"
	"k8scommerce/services/api/admin/internal/helpers"
	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
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

func (l *GetProductByIdLogic) GetProductById(req types.GetProductByIdRequest) (resp *types.Product, err error) {
	resp = &types.Product{}

	response, err := l.svcCtx.CatalogRpc.GetProductById(l.ctx, &catalogclient.GetProductByIdRequest{
		Id:      req.Id,
		StoreId: l.ctx.Value(types.StoreKey).(int64),
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	utils.TransformObj(response.Product, &resp)

	// format the currency to the locale and language
	for x := 0; x < len(resp.Variants); x++ {
		if resp.Variants[x].Price != (types.Price{}) {
			helpers.ConvertOutgoingPrices(l.ctx, &resp.Variants[x].Price)
		}
	}

	return resp, err
}
