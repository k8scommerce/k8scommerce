package logic

import (
	"context"
	"encoding/json"

	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
	"k8scommerce/services/rpc/product/productclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductByIdRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductByIdRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductByIdRequestLogic {
	return GetProductByIdRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductByIdRequestLogic) GetProductByIdRequest(req types.GetProductByIdRequest) (*types.Product, error) {
	getOneBySkuResponse, err := l.svcCtx.ProductRpc.GetProductById(l.ctx, &productclient.GetProductByIdRequest{
		Id: req.Id,
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
	return res, err
}
