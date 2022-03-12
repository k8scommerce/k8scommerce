package categories

import (
	"context"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCategoryByIdLogic {
	return GetCategoryByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryByIdLogic) GetCategoryById(req types.GetCategoryByIdRequest) (resp *types.Category, err error) {
	// todo: add your logic here and delete this line

	return
}
