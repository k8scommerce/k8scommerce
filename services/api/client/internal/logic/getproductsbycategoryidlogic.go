package logic

import (
	"context"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductsByCategoryIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductsByCategoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductsByCategoryIdLogic {
	return GetProductsByCategoryIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductsByCategoryIdLogic) GetProductsByCategoryId(req types.GetProductsByCategoryIdRequest) (resp *types.GetProductsByCategoryIdResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
