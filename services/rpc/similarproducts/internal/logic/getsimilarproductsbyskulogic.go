package logic

import (
	"context"
	"time"

	"k8scommerce/internal/gcache"
	"k8scommerce/internal/groupctx"
	"k8scommerce/services/rpc/similarproducts/internal/svc"
	"k8scommerce/services/rpc/similarproducts/pb/similarproducts"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

const Group_GetSimilarProductsBySku = "GetSimilarProductsBySku"

var Group_GetSimilarProductsBySkuKey = func(storeId int64, sku string) string {
	return gcache.ToKey(Group_GetSimilarProductsBySku, storeId, sku)
}

type GetSimilarProductsBySkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSimilarProductsBySkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSimilarProductsBySkuLogic {
	return &GetSimilarProductsBySkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSimilarProductsBySkuLogic) GetSimilarProductsBySku(in *similarproducts.GetSimilarProductsBySkuRequest) (*similarproducts.GetSimilarProductsBySkuResponse, error) {
	l.ctx = groupctx.SetStoreId(l.ctx, in.StoreId)
	l.ctx = groupctx.SetProductSku(l.ctx, in.Sku)
	res := &similarproducts.GetSimilarProductsBySkuResponse{}
	err := l.cache().Get(l.ctx, Group_GetSimilarProductsBySkuKey(in.StoreId, in.Sku), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetSimilarProductsBySkuLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetSimilarProductsBySku, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			// found, err := l.svcCtx.Repo.Category().GetSimilarProductsBySku(
			// 	groupctx.GetStoreId(ctx),
			// )
			// if err != nil {
			// 	logx.Infof("error: %s", err)
			// }

			// cats := []*catalog.Category{}

			// if found != nil {
			// 	for _, f := range found.Categories {
			// 		cat := catalog.Category{}
			// 		convert.ModelCategoryToProtoCategory(&f, &cat)
			// 		cats = append(cats, &cat)
			// 	}
			// }

			// Set the groupcache to expire after 24 hours
			if err := dest.SetProto(&similarproducts.GetSimilarProductsBySkuResponse{}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
