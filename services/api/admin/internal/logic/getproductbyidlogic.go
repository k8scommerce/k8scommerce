package logic

import (
	"context"

	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return
}
