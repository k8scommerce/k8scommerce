package logic

import (
	"context"
	"time"

	"k8scommerce/internal/convert"
	"k8scommerce/internal/gcache"
	"k8scommerce/internal/groupctx"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

const Group_GetProductBySku = "GetProductBySku"

var Group_GetProductBySkuKey = func(storeId int64, sku string) string {
	return gcache.ToKey(Group_GetProductBySku, storeId, sku)
}

type GetProductBySkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductBySkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductBySkuLogic {
	return &GetProductBySkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductBySkuLogic) GetProductBySku(in *catalog.GetProductBySkuRequest) (*catalog.GetProductBySkuResponse, error) {
	l.ctx = groupctx.SetStoreId(l.ctx, in.StoreId)
	l.ctx = groupctx.SetProductSku(l.ctx, in.Sku)
	res := &catalog.GetProductBySkuResponse{}
	err := l.cache().Get(l.ctx, Group_GetProductBySkuKey(in.StoreId, in.Sku), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetProductBySkuLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetAllCategories, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.Product().GetProductBySku(
				groupctx.GetStoreId(ctx),
				groupctx.GetProductSku(ctx),
			)
			if err != nil {
				logx.Infof("error: %s", err)
				return err
			}

			prod := catalog.Product{}
			if found != nil {
				convert.ModelProductToProtoProduct(&found.Product, &found.Variants, &found.Prices, &prod)

				for _, pair := range found.Categories {
					prod.Categories = append(prod.Categories, &catalog.CategoryPair{
						Slug: pair.Slug,
						Name: pair.Name,
					})
				}
			}

			// Set the groupcache to expire after 24 hours
			if err := dest.SetProto(&catalog.GetProductBySkuResponse{
				Product: &prod,
			}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
