package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(in *catalog.DeleteCategoryRequest) (*catalog.DeleteCategoryResponse, error) {
	_, err := l.svcCtx.Repo.Category().GetCategoryById(in.Id)
	if err != nil {
		return &catalog.DeleteCategoryResponse{}, err
	}

	// delete the Category
	if err := l.svcCtx.Repo.Category().Delete(in.Id); err != nil {
		return &catalog.DeleteCategoryResponse{
			// StatusCode:    http.StatusExpectationFailed,
			// StatusMessage: err.Error(),
		}, err
	}

	// invalidate the cache for this record
	{
		l.svcCtx.Cache.Delete(l.ctx, Group_GetCategoryById, Group_GetCategoryByIdKey(in.Id))
		l.svcCtx.Cache.Delete(l.ctx, Group_GetAllCategories, Group_GetAllCategoriesKey(in.StoreId))
	}

	// the response struct
	return &catalog.DeleteCategoryResponse{
		// StatusCode:    http.StatusOK,
		// StatusMessage: "",
	}, nil
}
