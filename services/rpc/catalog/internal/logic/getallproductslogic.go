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

const Group_GetAllProducts = "GetAllProducts"

var Group_GetAllProductsKey = func(storeId, currentPage, pageSize int64, filter string) string {
	return gcache.ToKey(Group_GetAllProducts, storeId, currentPage, pageSize, filter)
}

type GetAllProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllProductsLogic {
	return &GetAllProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllProductsLogic) GetAllProducts(in *catalog.GetAllProductsRequest) (*catalog.GetAllProductsResponse, error) {
	l.ctx = groupctx.SetStoreId(l.ctx, in.StoreId)
	l.ctx = groupctx.SetCurrentPage(l.ctx, in.CurrentPage)
	l.ctx = groupctx.SetPageSize(l.ctx, in.PageSize)
	l.ctx = groupctx.SetFilter(l.ctx, in.Filter)
	res := &catalog.GetAllProductsResponse{}
	err := l.cache().Get(l.ctx, Group_GetAllProductsKey(in.StoreId, in.CurrentPage, in.PageSize, in.Filter), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetAllProductsLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetAllCategories, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.Product().GetAllProducts(
				groupctx.GetStoreId(ctx),
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
					defaultImage := []*models.Asset{}
					defaultImage = append(defaultImage, &f.Asset)

					convertedImages := convert.ModelAssetToProtoAsset(defaultImage)
					if len(convertedImages) > 0 {
						prod.DefaultImage = convertedImages[0]
					}

					convert.ModelProductToProtoProduct(&f.Product, &[]models.Variant{
						f.Variant,
					}, &[]models.Price{
						f.Price,
					}, &prod)
					prods = append(prods, &prod)
				}
			}

			// Set the groupcache to expire after 24 hours
			if err := dest.SetProto(&catalog.GetAllProductsResponse{
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
