package logic

import (
	"context"
	"time"

	"k8scommerce/internal/convert"
	"k8scommerce/internal/gcache"
	"k8scommerce/internal/groupctx"
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

const Group_GetProductsByCategorySlug = " GetProductsByCategorySlug"

var Group_GetProductsByCategorySlugKey = func(storeId int64, categorySlug string, currentPage, pageSize int64, filter string) string {
	return gcache.ToKey(Group_GetProductsByCategorySlug, storeId, categorySlug, currentPage, pageSize, filter)
}

type GetProductsByCategorySlugLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductsByCategorySlugLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductsByCategorySlugLogic {
	return &GetProductsByCategorySlugLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductsByCategorySlugLogic) GetProductsByCategorySlug(in *catalog.GetProductsByCategorySlugRequest) (*catalog.GetProductsByCategorySlugResponse, error) {
	l.ctx = groupctx.SetStoreId(l.ctx, in.StoreId)
	l.ctx = groupctx.SetCategorySlug(l.ctx, in.CategorySlug)
	l.ctx = groupctx.SetCurrentPage(l.ctx, in.CurrentPage)
	l.ctx = groupctx.SetPageSize(l.ctx, in.PageSize)
	l.ctx = groupctx.SetFilter(l.ctx, in.Filter)
	res := &catalog.GetProductsByCategorySlugResponse{}
	err := l.cache().Get(l.ctx, Group_GetProductsByCategorySlugKey(in.StoreId, in.CategorySlug, in.CurrentPage, in.PageSize, in.Filter), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetProductsByCategorySlugLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetAllCategories, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.Product().GetProductsByCategorySlug(
				groupctx.GetStoreId(ctx),
				groupctx.GetCategorySlug(ctx),
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
			if err := dest.SetProto(&catalog.GetProductsByCategorySlugResponse{
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
