package logic

import (
	"context"
	"time"

	"k8scommerce/internal/gcache"
	"k8scommerce/services/rpc/store/internal/svc"
	"k8scommerce/services/rpc/store/pb/store"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

const Group_GetAllStores = "GetAllStores"

var Group_GetAllStoresKey = func() string {
	return gcache.ToKey(Group_GetAllStores)
}

type GetAllStoresLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllStoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllStoresLogic {
	return &GetAllStoresLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllStoresLogic) GetAllStores(in *store.GetAllStoresRequest) (*store.GetAllStoresResponse, error) {
	// l.ctx = groupctx.SetCurrentPage(l.ctx, in.CurrentPage)
	// l.ctx = groupctx.SetPageSize(l.ctx, in.PageSize)
	// l.ctx = groupctx.SetSortOn(l.ctx, in.SortOn)
	res := &store.GetAllStoresResponse{}
	err := l.cache().Get(l.ctx, Group_GetAllStoresKey(), groupcache.ProtoSink(res))
	return res, err

}

func (l *GetAllStoresLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetAllStores, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			// found, err := l.svcCtx.Repo.Category().GetAllStores(
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
			if err := dest.SetProto(&store.GetAllStoresResponse{}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
