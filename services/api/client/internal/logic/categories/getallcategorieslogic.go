package categories

import (
	"context"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllCategoriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllCategoriesLogic {
	return GetAllCategoriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllCategoriesLogic) GetAllCategories() (resp *types.GetAllCategoriesResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
