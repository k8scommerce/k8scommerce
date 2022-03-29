package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/k8scommerce/k8scommerce/internal/models"
	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCategoryLogic) CreateCategory(in *catalog.CreateCategoryRequest) (*catalog.CreateCategoryResponse, error) {
	prod := models.Category{}
	utils.TransformObj(in.Category, prod)
	if err := l.svcCtx.Repo.Category().Create(&prod); err != nil {
		logx.Infof("error: %s", err)
		return &catalog.CreateCategoryResponse{
			Category: nil,
		}, err
	}

	// the output object
	out := &catalog.Category{}
	utils.TransformObj(prod, &out)

	{
		l.svcCtx.Cache.DestroyGroup(Group_GetAllCategories)
	}
	// the response struct
	return &catalog.CreateCategoryResponse{
		Category: out,
	}, nil
}
