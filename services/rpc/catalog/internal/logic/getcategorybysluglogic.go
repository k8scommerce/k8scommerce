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

const Group_GetCategoryBySlug = "GetCategoryBySlug"

var Group_GetCategoryBySlugKey = func(storeId int64, slug string) string {
	return gcache.ToKey(Group_GetCategoryBySlug, storeId, slug)
}

type GetCategoryBySlugLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryBySlugLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryBySlugLogic {
	return &GetCategoryBySlugLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryBySlugLogic) GetCategoryBySlug(in *catalog.GetCategoryBySlugRequest) (*catalog.GetCategoryBySlugResponse, error) {
	l.ctx = groupctx.SetStoreId(l.ctx, in.StoreId)
	l.ctx = groupctx.SetCategorySlug(l.ctx, in.Slug)
	res := &catalog.GetCategoryBySlugResponse{}
	err := l.cache().Get(l.ctx, Group_GetCategoryBySlugKey(in.StoreId, in.Slug), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetCategoryBySlugLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetAllCategories, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.Category().GetCategoryBySlug(
				groupctx.GetStoreId(ctx),
				groupctx.GetCategorySlug(ctx),
			)
			if err != nil {
				logx.Infof("error: %s", err)
				return err
			}

			cat := catalog.Category{}
			if found != nil {
				convert.ModelCategoryToProtoCategory(found, &cat)
			}

			// Set the groupcache to expire after 24 hours
			if err := dest.SetProto(&catalog.GetCategoryBySlugResponse{
				Category: &cat,
			}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
