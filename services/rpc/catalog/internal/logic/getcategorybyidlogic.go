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

const Group_GetCategoryById = "GetCategoryById"

var Group_GetCategoryByIdKey = func(categoryId int64) string {
	return gcache.ToKey(Group_GetCategoryById, categoryId)
}

type GetCategoryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryByIdLogic {
	return &GetCategoryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryByIdLogic) GetCategoryById(in *catalog.GetCategoryByIdRequest) (*catalog.GetCategoryByIdResponse, error) {
	l.ctx = groupctx.SetCategoryId(l.ctx, in.Id)
	res := &catalog.GetCategoryByIdResponse{}
	err := l.cache().Get(l.ctx, Group_GetCategoryByIdKey(in.Id), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetCategoryByIdLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetAllCategories, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.Category().GetCategoryById(
				groupctx.GetCategoryId(ctx),
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
			if err := dest.SetProto(&catalog.GetCategoryByIdResponse{
				Category: &cat,
			}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
