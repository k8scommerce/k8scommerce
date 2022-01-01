package logic

import (
	"context"
	"net/http"
	"strconv"
	"sync"

	"github.com/k8s-commerce/k8s-commerce/pkg/models"
	"github.com/k8s-commerce/k8s-commerce/pkg/utils"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/product/internal/svc"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/product/product"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *product.UpdateProductRequest) (*product.UpdateProductResponse, error) {
	found, err := l.svcCtx.Repo.Product().GetProductById(in.Id)
	if err != nil {
		return &product.UpdateProductResponse{
			StatusCode:    http.StatusExpectationFailed,
			StatusMessage: err.Error(),
		}, nil
	}

	prod := models.Product{}
	utils.TransformObj(in.Product, prod)
	prod.ID = found.Product.ID // make sure we're updating the correct id
	if err := l.svcCtx.Repo.Product().Update(&prod); err != nil {
		logx.Infof("error: %s", err)
		return &product.UpdateProductResponse{
			Product:       nil,
			StatusCode:    http.StatusExpectationFailed,
			StatusMessage: err.Error(),
		}, nil
	}

	// get the sku from the primary variant
	var sku string
	for _, variant := range found.Variants {
		if variant.IsDefault {
			sku = variant.Sku
		}
	}

	// invalidate the cache for this record
	{
		if entryGetProductByIdLogic != nil {
			l.mu.Lock()
			entryGetProductByIdLogic.galaxy.Remove(l.ctx, strconv.Itoa(int(in.Id)))
			l.mu.Unlock()
		}
		if entryGetProductBySkuLogic != nil {
			l.mu.Lock()
			entryGetProductBySkuLogic.galaxy.Remove(l.ctx, sku)
			l.mu.Unlock()
		}
	}

	// the output object
	out := &product.Product{}
	utils.TransformObj(prod, &out)

	// the response struct
	return &product.UpdateProductResponse{
		Product:       out,
		StatusCode:    http.StatusOK,
		StatusMessage: "",
	}, nil
}
