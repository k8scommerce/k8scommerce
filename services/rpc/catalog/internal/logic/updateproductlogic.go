package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/k8scommerce/k8scommerce/internal/models"
	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *catalog.UpdateProductRequest) (*catalog.UpdateProductResponse, error) {
	found, err := l.svcCtx.Repo.Product().GetProductById(in.Id)
	if err != nil {
		return &catalog.UpdateProductResponse{}, err
	}

	prod := models.Product{}
	utils.TransformObj(in.Product, prod)
	prod.ID = found.Product.ID // make sure we're updating the correct id
	if err := l.svcCtx.Repo.Product().Update(&prod); err != nil {
		logx.Infof("error: %s", err)
		return &catalog.UpdateProductResponse{
			Product: nil,
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
		l.svcCtx.Cache.Delete(l.ctx, Group_GetProductById, Group_GetProductByIdKey(prod.ID))
		l.svcCtx.Cache.Delete(l.ctx, Group_GetProductBySku, Group_GetProductBySkuKey(in.StoreId, sku))
		l.svcCtx.Cache.Delete(l.ctx, Group_GetProductBySlug, Group_GetProductBySlugKey(in.StoreId, prod.Slug))

		l.svcCtx.Cache.DestroyGroup(Group_GetAllProducts)
		l.svcCtx.Cache.DestroyGroup(Group_GetProductsByCategoryId)
	}

	// the output object
	out := &catalog.Product{}
	utils.TransformObj(prod, &out)

	// the response struct
	return &catalog.UpdateProductResponse{
		Product: out,
	}, err
}
