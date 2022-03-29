package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/k8scommerce/k8scommerce/internal/models"
	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductLogic) CreateProduct(in *catalog.CreateProductRequest) (*catalog.CreateProductResponse, error) {
	prod := models.Product{}
	utils.TransformObj(in.Product, prod)
	if err := l.svcCtx.Repo.Product().Create(&prod); err != nil {
		logx.Infof("error: %s", err)
		return &catalog.CreateProductResponse{
			Product: nil,
		}, nil
	}

	// the output object
	out := &catalog.Product{}
	utils.TransformObj(prod, &out)

	// the response struct
	return &catalog.CreateProductResponse{
		Product: out,
	}, nil
}
