package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"

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
	// todo: add your logic here and delete this line

	return &types.Product{}, nil
}