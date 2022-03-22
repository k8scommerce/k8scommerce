package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"k8scommerce/internal/convert"
	"k8scommerce/internal/galaxyctx"
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/zeromicro/go-zero/core/logx"
)

type galaxyGetProductsByCategorySlugLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetProductsByCategorySlugLogic *galaxyGetProductsByCategorySlugLogicHelper

type GetProductsByCategorySlugLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetProductsByCategorySlugLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetProductsByCategorySlugLogic {
	return &GetProductsByCategorySlugLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetProductsByCategorySlugLogic) GetProductsByCategorySlug(in *catalog.GetProductsByCategorySlugRequest) (*catalog.GetProductsByCategorySlugResponse, error) {

	// caching goes logic here
	if entryGetProductsByCategorySlugLogic == nil {
		l.mu.Lock()
		entryGetProductsByCategorySlugLogic = &galaxyGetProductsByCategorySlugLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetProductsByCategorySlugLogic.once.Do(func() {

		// register the galaxy one time
		entryGetProductsByCategorySlugLogic.galaxy = gcache.RegisterGalaxyFunc("GetProductsByCategorySlug", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {

				found, err := l.svcCtx.Repo.Product().GetProductsByCategorySlug(
					galaxyctx.GetStoreId(ctx),
					galaxyctx.GetCategorySlug(ctx),
					galaxyctx.GetCurrentPage(ctx),
					galaxyctx.GetPageSize(ctx),
					galaxyctx.GetFilter(ctx),
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

				// the response struct
				item := &catalog.GetProductsByCategorySlugResponse{
					Products:     prods,
					TotalRecords: totalRecords,
					TotalPages:   totalPages,
				}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	codec := &galaxycache.ByteCodec{}

	l.ctx = galaxyctx.SetStoreId(l.ctx, in.StoreId)
	l.ctx = galaxyctx.SetCategorySlug(l.ctx, in.CategorySlug)
	l.ctx = galaxyctx.SetCurrentPage(l.ctx, in.CurrentPage)
	l.ctx = galaxyctx.SetPageSize(l.ctx, in.PageSize)
	l.ctx = galaxyctx.SetFilter(l.ctx, in.Filter)

	key := fmt.Sprintf("%d|%s|%d|%d|%s", in.StoreId, in.CategorySlug, in.CurrentPage, in.PageSize, in.Filter)
	entryGetProductsByCategorySlugLogic.galaxy.Get(l.ctx, key, codec)
	b, err := codec.MarshalBinary()
	if err != nil {
		return nil, err
	}
	res := &catalog.GetProductsByCategorySlugResponse{}
	err = json.Unmarshal(b, res)

	return res, err
}
