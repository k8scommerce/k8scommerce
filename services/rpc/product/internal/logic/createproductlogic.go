package logic

import (
	"context"
	"net/http"

	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/product/internal/svc"
	"k8scommerce/services/rpc/product/pb/product"

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

func (l *CreateProductLogic) CreateProduct(in *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	prod := models.Product{}
	utils.TransformObj(in.Product, prod)
	if err := l.svcCtx.Repo.Product().Create(&prod); err != nil {
		logx.Infof("error: %s", err)
		return &product.CreateProductResponse{
			Product:       nil,
			StatusCode:    http.StatusExpectationFailed,
			StatusMessage: err.Error(),
		}, nil
	}

	// the output object
	out := &product.Product{}
	utils.TransformObj(prod, &out)

	// the response struct
	return &product.CreateProductResponse{
		Product:       out,
		StatusCode:    http.StatusOK,
		StatusMessage: "",
	}, nil
}
