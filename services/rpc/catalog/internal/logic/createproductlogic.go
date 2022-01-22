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

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *CreateProductLogic) CreateProduct(in *catalog.CreateProductRequest) (*catalog.CreateProductResponse, error) {
	prod := models.Product{}
	utils.TransformObj(in.Product, prod)
	if err := l.svcCtx.Repo.Product().Create(&prod); err != nil {
		logx.Infof("error: %s", err)
		return &catalog.CreateProductResponse{
			Product: nil,
			// StatusCode:    http.StatusExpectationFailed,
			// StatusMessage: err.Error(),
		}, nil
	}

	// the output object
	out := &catalog.Product{}
	utils.TransformObj(prod, &out)

	// the response struct
	return &catalog.CreateProductResponse{
		Product: out,
		// StatusCode:    http.StatusOK,
		// StatusMessage: "",
	}, nil
}
