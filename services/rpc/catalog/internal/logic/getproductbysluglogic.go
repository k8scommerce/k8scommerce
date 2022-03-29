package logic

import (
	"context"
	"time"

	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/k8scommerce/k8scommerce/internal/convert"
	"github.com/k8scommerce/k8scommerce/internal/gcache"
	"github.com/k8scommerce/k8scommerce/internal/groupctx"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

const Group_GetProductBySlug = "GetProductBySlug"

var Group_GetProductBySlugKey = func(storeId int64, slug string) string {
	return gcache.ToKey(Group_GetProductBySlug, storeId, slug)
}

type GetProductBySlugLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductBySlugLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductBySlugLogic {
	return &GetProductBySlugLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductBySlugLogic) GetProductBySlug(in *catalog.GetProductBySlugRequest) (*catalog.GetProductBySlugResponse, error) {
	l.ctx = groupctx.SetStoreId(l.ctx, in.StoreId)
	l.ctx = groupctx.SetProductSku(l.ctx, in.Slug)
	res := &catalog.GetProductBySlugResponse{}
	err := l.cache().Get(l.ctx, Group_GetProductBySlugKey(in.StoreId, in.Slug), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetProductBySlugLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetAllCategories, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.Product().GetProductBySlug(
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
			if err := dest.SetProto(&catalog.GetProductBySlugResponse{
				Product: &prod,
			}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
