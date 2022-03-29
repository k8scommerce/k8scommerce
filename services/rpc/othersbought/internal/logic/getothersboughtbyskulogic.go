package logic

import (
	"context"
	"time"

	"github.com/k8scommerce/k8scommerce/services/rpc/othersbought/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/othersbought/pb/othersbought"

	"github.com/k8scommerce/k8scommerce/internal/gcache"
	"github.com/k8scommerce/k8scommerce/internal/groupctx"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

const Group_GetOthersBoughtBySku = "GetOthersBoughtBySku"

var Group_GetOthersBoughtBySkuKey = func(storeId int64, sku string) string {
	return gcache.ToKey(Group_GetOthersBoughtBySku, storeId, sku)
}

type GetOthersBoughtBySkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOthersBoughtBySkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOthersBoughtBySkuLogic {
	return &GetOthersBoughtBySkuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOthersBoughtBySkuLogic) GetOthersBoughtBySku(in *othersbought.GetOthersBoughtBySkuRequest) (*othersbought.GetOthersBoughtBySkuResponse, error) {
	l.ctx = groupctx.SetStoreId(l.ctx, in.StoreId)
	l.ctx = groupctx.SetProductSku(l.ctx, in.Sku)
	res := &othersbought.GetOthersBoughtBySkuResponse{}
	err := l.cache().Get(l.ctx, Group_GetOthersBoughtBySkuKey(in.StoreId, in.Sku), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetOthersBoughtBySkuLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetOthersBoughtBySku, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			// found, err := l.svcCtx.Repo.Category().GetAllCategories(
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
			if err := dest.SetProto(&othersbought.GetOthersBoughtBySkuResponse{}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
