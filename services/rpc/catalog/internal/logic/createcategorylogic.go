package logic

import (
	"context"
	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *CreateCategoryLogic) CreateCategory(in *catalog.CreateCategoryRequest) (*catalog.CreateCategoryResponse, error) {
	prod := models.Category{}
	utils.TransformObj(in.Category, prod)
	if err := l.svcCtx.Repo.Category().Create(&prod); err != nil {
		logx.Infof("error: %s", err)
		return &catalog.CreateCategoryResponse{
			Category: nil,
			// StatusCode:    http.StatusExpectationFailed,
			// StatusMessage: err.Error(),
		}, err
	}

	// the output object
	out := &catalog.Category{}
	utils.TransformObj(prod, &out)

	// the response struct
	return &catalog.CreateCategoryResponse{
		Category: out,
		// StatusCode:    http.StatusOK,
		// StatusMessage: "",
	}, nil
}
