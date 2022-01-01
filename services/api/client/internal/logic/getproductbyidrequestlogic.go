package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"

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
	// todo: add your logic here and delete this line

	return &types.Product{}, nil
}