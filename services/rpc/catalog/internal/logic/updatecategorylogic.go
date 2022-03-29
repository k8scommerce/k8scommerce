package logic

import (
	"context"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *catalog.UpdateCategoryRequest) (*catalog.UpdateCategoryResponse, error) {
	found, err := l.svcCtx.Repo.Category().GetCategoryById(in.Id)
	if err != nil {
		return &catalog.UpdateCategoryResponse{}, err
	}

	if err := l.svcCtx.Repo.Category().Update(found); err != nil {
		logx.Infof("error: %s", err)
		return &catalog.UpdateCategoryResponse{
			Category: nil,
		}, err
	}

	// invalidate the cache for this record
	{
		l.svcCtx.Cache.Delete(l.ctx, Group_GetCategoryById, Group_GetCategoryByIdKey(in.Id))
		l.svcCtx.Cache.Delete(l.ctx, Group_GetAllCategories, Group_GetAllCategoriesKey(in.StoreId))
		l.svcCtx.Cache.DestroyGroup(Group_GetProductsByCategoryId)
	}

	// the output object
	out := &catalog.Category{}
	utils.TransformObj(found, &out)

	// the response struct
	return &catalog.UpdateCategoryResponse{
		Category: out,
	}, err

}
