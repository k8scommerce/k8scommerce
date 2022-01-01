package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductBySlugRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductBySlugRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductBySlugRequestLogic {
	return GetProductBySlugRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductBySlugRequestLogic) GetProductBySlugRequest(req types.GetProductBySlugRequest) (*types.Product, error) {
	// todo: add your logic here and delete this line

	return &types.Product{}, nil
}
