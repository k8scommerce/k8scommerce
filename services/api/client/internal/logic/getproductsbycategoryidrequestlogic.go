package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductsByCategoryIdRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductsByCategoryIdRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductsByCategoryIdRequestLogic {
	return GetProductsByCategoryIdRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductsByCategoryIdRequestLogic) GetProductsByCategoryIdRequest(req types.GetProductsByCategoryIdRequest) (*types.GetProductsByCategoryIdResponse, error) {
	// todo: add your logic here and delete this line

	return &types.GetProductsByCategoryIdResponse{}, nil
}
