package logic

import (
	"context"

	"net/http"
	"strconv"
	"sync"

	"github.com/k8s-commerce/k8s-commerce/services/rpc/product/internal/svc"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/product/product"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *DeleteProductLogic) DeleteProduct(in *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	prod, err := l.svcCtx.Repo.Product().GetProductById(in.Id)
	if err != nil {
		return &product.DeleteProductResponse{
			StatusCode:    http.StatusExpectationFailed,
			StatusMessage: err.Error(),
		}, nil
	}

	// get the sku from the primary variant
	var sku string
	for _, variant := range prod.Variants {
		if variant.IsDefault {
			sku = variant.Sku
		}
	}

	// delete the product
	if err := l.svcCtx.Repo.Product().Delete(in.Id); err != nil {
		return &product.DeleteProductResponse{
			StatusCode:    http.StatusExpectationFailed,
			StatusMessage: err.Error(),
		}, nil
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

	// the response struct
	return &product.DeleteProductResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "",
	}, nil
}
