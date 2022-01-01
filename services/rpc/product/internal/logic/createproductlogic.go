package logic

import (
	"context"
	"net/http"

	"github.com/k8s-commerce/k8s-commerce/pkg/models"
	"github.com/k8s-commerce/k8s-commerce/pkg/utils"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/product/internal/svc"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/product/product"

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
