package Products

import (
	"context"

	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
