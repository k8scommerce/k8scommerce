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

const Group_GetAllCategories = "GetAllCategories"

var Group_GetAllCategoriesKey = func(storeId int64) string {
	return gcache.ToKey(Group_GetAllCategories, storeId)
}

type GetAllCategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllCategoriesLogic {
	return &GetAllCategoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  categories
func (l *GetAllCategoriesLogic) GetAllCategories(in *catalog.GetAllCategoriesRequest) (*catalog.GetAllCategoriesResponse, error) {
	l.ctx = groupctx.SetStoreId(l.ctx, in.StoreId)
	res := &catalog.GetAllCategoriesResponse{}
	err := l.cache().Get(l.ctx, Group_GetAllCategoriesKey(in.StoreId), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetAllCategoriesLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetAllCategories, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.Category().GetAllCategories(
				groupctx.GetStoreId(ctx),
			)
			if err != nil {
				logx.Infof("error: %s", err)
				return err
			}

			cats := []*catalog.Category{}

			if found != nil {
				for _, f := range found.Categories {
					cat := catalog.Category{}
					convert.ModelCategoryToProtoCategory(&f, &cat)
					cats = append(cats, &cat)
				}
			}

			// Set the groupcache to expire after 24 hours
			if err := dest.SetProto(&catalog.GetAllCategoriesResponse{
				Categories: cats,
			}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
