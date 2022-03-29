package logic

import (
	"context"
	"time"

	"github.com/k8scommerce/k8scommerce/services/rpc/store/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/store/pb/store"

	"github.com/k8scommerce/k8scommerce/internal/convert"
	"github.com/k8scommerce/k8scommerce/internal/gcache"
	"github.com/k8scommerce/k8scommerce/internal/groupctx"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

const Group_GetStoreById = "GetStoreById"

var Group_GetStoreByIdKey = func(storeId int64) string {
	return gcache.ToKey(Group_GetStoreById, storeId)
}

type GetStoreByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStoreByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStoreByIdLogic {
	return &GetStoreByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStoreByIdLogic) GetStoreById(in *store.GetStoreByIdRequest) (*store.GetStoreByIdResponse, error) {
	l.ctx = groupctx.SetStoreId(l.ctx, in.Id)
	res := &store.GetStoreByIdResponse{}
	err := l.cache().Get(l.ctx, Group_GetStoreByIdKey(in.Id), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetStoreByIdLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetStoreById, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.Store().GetStoreById(groupctx.GetStoreId(ctx))
			if err != nil {
				logx.Infof("error: %s", err)
				return err
			}

			s := &store.Store{}
			if found != nil {
				convert.ModelStoreToProtoStore(found, s)
			}

			// Set the groupcache to expire after 24 hours
			if err := dest.SetProto(&store.GetStoreByIdResponse{
				Store: s,
			}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
