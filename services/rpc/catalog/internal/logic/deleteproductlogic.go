package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductLogic) DeleteProduct(in *catalog.DeleteProductRequest) (*catalog.DeleteProductResponse, error) {
	prod, err := l.svcCtx.Repo.Product().GetProductById(in.Id)
	if err != nil {
		return &catalog.DeleteProductResponse{}, err
	}

	if prod != nil {
		// get the sku from the primary variant
		var sku string
		for _, variant := range prod.Variants {
			if variant.IsDefault {
				sku = variant.Sku
			}
		}

		// delete the product
		if err := l.svcCtx.Repo.Product().Delete(prod.Product.ID); err != nil {
			return &catalog.DeleteProductResponse{}, err
		}

		// invalidate the cache for this record
		{
			l.svcCtx.Cache.Delete(l.ctx, Group_GetProductById, Group_GetProductByIdKey(prod.Product.ID))
			l.svcCtx.Cache.Delete(l.ctx, Group_GetProductBySku, Group_GetProductBySkuKey(in.StoreId, sku))
			l.svcCtx.Cache.Delete(l.ctx, Group_GetProductBySlug, Group_GetProductBySlugKey(in.StoreId, prod.Product.Slug))

			l.svcCtx.Cache.DestroyGroup(Group_GetAllProducts)
			l.svcCtx.Cache.DestroyGroup(Group_GetProductsByCategoryId)
		}
	}

	// the response struct
	return &catalog.DeleteProductResponse{}, nil
}
