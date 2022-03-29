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

const Group_GetWarehouseById = "GetWarehouseById"

var Group_GetWarehouseByIdKey = func(storeId int64) string {
	return gcache.ToKey(Group_GetWarehouseById, storeId)
}

type GetWarehouseByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWarehouseByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWarehouseByIdLogic {
	return &GetWarehouseByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWarehouseByIdLogic) GetWarehouseById(in *warehouse.GetWarehouseByIdRequest) (*warehouse.GetWarehouseByIdResponse, error) {
	l.ctx = groupctx.SetWarehouseId(l.ctx, in.Id)
	res := &warehouse.GetWarehouseByIdResponse{}
	err := l.cache().Get(l.ctx, Group_GetWarehouseByIdKey(in.Id), groupcache.ProtoSink(res))
	return res, err

}

func (l *GetWarehouseByIdLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetWarehouseById, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			// found, err := l.svcCtx.Repo.Category().GetWarehouseById(
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
			if err := dest.SetProto(&warehouse.GetWarehouseByIdResponse{}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
