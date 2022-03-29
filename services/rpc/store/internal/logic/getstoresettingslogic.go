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

const Group_GetStoreSettings = "GetStoreSettings"

var Group_GetStoreSettingsKey = func(storeId int64) string {
	return gcache.ToKey(Group_GetStoreSettings, storeId)
}

type GetStoreSettingsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStoreSettingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStoreSettingsLogic {
	return &GetStoreSettingsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStoreSettingsLogic) GetStoreSettings(in *store.GetStoreSettingRequest) (*store.GetStoreSettingResponse, error) {

	l.ctx = groupctx.SetStoreId(l.ctx, in.StoreId)
	res := &store.GetStoreSettingResponse{}
	err := l.cache().Get(l.ctx, Group_GetStoreSettingsKey(in.StoreId), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetStoreSettingsLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetStoreSettings, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.StoreSetting().GetStoreSettingById(groupctx.GetStoreId(ctx))
			if err != nil {
				logx.Infof("error: %s", err)
				return err
			}

			setting := &store.StoreSetting{}
			if found != nil {
				convert.ModelStoreSettingToProtoStoreSetting(found, setting)
			}

			// Set the groupcache to expire after 24 hours
			if err := dest.SetProto(&store.GetStoreSettingResponse{
				Setting: setting,
			}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
