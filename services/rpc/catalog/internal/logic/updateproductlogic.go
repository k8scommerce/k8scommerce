package logic

import (
	"context"
	"strconv"
	"sync"

	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"

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

func (l *UpdateProductLogic) UpdateProduct(in *catalog.UpdateProductRequest) (*catalog.UpdateProductResponse, error) {
	found, err := l.svcCtx.Repo.Product().GetProductById(in.Id)
	if err != nil {
		return &catalog.UpdateProductResponse{
			// StatusCode:    http.StatusExpectationFailed,
			// StatusMessage: err.Error(),
		}, err
	}

	prod := models.Product{}
	utils.TransformObj(in.Product, prod)
	prod.ID = found.Product.ID // make sure we're updating the correct id
	if err := l.svcCtx.Repo.Product().Update(&prod); err != nil {
		logx.Infof("error: %s", err)
		return &catalog.UpdateProductResponse{
			Product: nil,
			// StatusCode:    http.StatusExpectationFailed,
			// StatusMessage: err.Error(),
		}, err
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
	out := &catalog.Product{}
	utils.TransformObj(prod, &out)

	// the response struct
	return &catalog.UpdateProductResponse{
		Product: out,
		// StatusCode:    http.StatusOK,
		// StatusMessage: "",
	}, err
}
