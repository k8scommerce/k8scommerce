package categories

import (
	"context"

	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCategoryBySlugLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryBySlugLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCategoryBySlugLogic {
	return GetCategoryBySlugLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryBySlugLogic) GetCategoryBySlug(req types.GetCategoryBySlugRequest) (resp *types.Category, err error) {
	// todo: add your logic here and delete this line

	return
}
