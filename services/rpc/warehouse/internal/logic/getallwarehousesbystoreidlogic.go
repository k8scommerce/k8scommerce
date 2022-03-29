package logic

import (
	"context"
	"time"

	"github.com/k8scommerce/k8scommerce/services/rpc/warehouse/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/warehouse/pb/warehouse"

	"github.com/k8scommerce/k8scommerce/internal/gcache"
	"github.com/k8scommerce/k8scommerce/internal/groupctx"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

const Group_GetAllWarehousesByStoreId = "GetAllWarehousesByStoreId"

var Group_GetAllWarehousesByStoreIdKey = func(storeId int64) string {
	return gcache.ToKey(Group_GetAllWarehousesByStoreId, storeId)
}

type GetAllWarehousesByStoreIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllWarehousesByStoreIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllWarehousesByStoreIdLogic {
	return &GetAllWarehousesByStoreIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllWarehousesByStoreIdLogic) GetAllWarehousesByStoreId(in *warehouse.GetAllWarehousesByStoreIdRequest) (*warehouse.GetAllWarehousesByStoreIdResponse, error) {
	l.ctx = groupctx.SetStoreId(l.ctx, in.StoreId)
	res := &warehouse.GetAllWarehousesByStoreIdResponse{}
	err := l.cache().Get(l.ctx, Group_GetAllWarehousesByStoreIdKey(in.StoreId), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetAllWarehousesByStoreIdLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetAllWarehousesByStoreId, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			// found, err := l.svcCtx.Repo.Category().GetAllWarehousesByStoreId(
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
			if err := dest.SetProto(&warehouse.GetAllWarehousesByStoreIdResponse{}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
