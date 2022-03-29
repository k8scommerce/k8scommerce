package logic

import (
	"context"
	"time"

	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/k8scommerce/k8scommerce/internal/convert"
	"github.com/k8scommerce/k8scommerce/internal/gcache"
	"github.com/k8scommerce/k8scommerce/internal/groupctx"
	"github.com/k8scommerce/k8scommerce/internal/models"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

const Group_GetProductsByCategoryId = " GetProductsByCategoryId"

var Group_GetProductsByCategoryIdKey = func(storeId, categoryId, currentPage, pageSize int64, filter string) string {
	return gcache.ToKey(Group_GetProductsByCategoryId, storeId, categoryId, currentPage, pageSize, filter)
}

type GetProductsByCategoryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductsByCategoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductsByCategoryIdLogic {
	return &GetProductsByCategoryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductsByCategoryIdLogic) GetProductsByCategoryId(in *catalog.GetProductsByCategoryIdRequest) (*catalog.GetProductsByCategoryIdResponse, error) {
	l.ctx = groupctx.SetStoreId(l.ctx, in.StoreId)
	l.ctx = groupctx.SetCategoryId(l.ctx, in.CategoryId)
	l.ctx = groupctx.SetCurrentPage(l.ctx, in.CurrentPage)
	l.ctx = groupctx.SetPageSize(l.ctx, in.PageSize)
	l.ctx = groupctx.SetFilter(l.ctx, in.Filter)
	res := &catalog.GetProductsByCategoryIdResponse{}
	err := l.cache().Get(l.ctx, Group_GetProductsByCategoryIdKey(in.StoreId, in.CategoryId, in.CurrentPage, in.PageSize, in.Filter), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetProductsByCategoryIdLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetAllCategories, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.Product().GetProductsByCategoryId(
				groupctx.GetStoreId(ctx),
				groupctx.GetCategoryId(ctx),
				groupctx.GetCurrentPage(ctx),
				groupctx.GetPageSize(ctx),
				groupctx.GetFilter(ctx),
			)
			if err != nil {
				logx.Infof("error: %s", err)
				return err
			}

			prods := []*catalog.Product{}

			var totalRecords int64 = 0
			var totalPages int64 = 0

			if found != nil {
				totalRecords = found.PagingStats.TotalRecords
				totalPages = found.PagingStats.TotalPages

				for _, f := range found.Results {
					prod := catalog.Product{}

					convert.ModelProductToProtoProduct(&f.Product, &[]models.Variant{
						f.Variant,
					}, &[]models.Price{
						f.Price,
					}, &prod)
					prods = append(prods, &prod)
				}
			}

			// Set the groupcache to expire after 24 hours
			if err := dest.SetProto(&catalog.GetProductsByCategoryIdResponse{
				Products:     prods,
				TotalRecords: totalRecords,
				TotalPages:   totalPages,
			}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
