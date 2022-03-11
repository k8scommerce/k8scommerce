package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"k8scommerce/internal/galaxyctx"
	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/internal/types"
	"k8scommerce/services/rpc/catalog/pb/catalog"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/zeromicro/go-zero/core/logx"
)

type galaxyGetAllProductsLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetAllProductsLogic *galaxyGetAllProductsLogicHelper

type GetAllProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetAllProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetAllProductsLogic {
	return &GetAllProductsLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetAllProductsLogic) GetAllProducts(in *catalog.GetAllProductsRequest) (*catalog.GetAllProductsResponse, error) {

	// caching goes logic here
	if entryGetAllProductsLogic == nil {
		l.mu.Lock()
		entryGetAllProductsLogic = &galaxyGetAllProductsLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetAllProductsLogic.once.Do(func() {
		// register the galaxy one time
		entryGetAllProductsLogic.galaxy = gcache.RegisterGalaxyFunc("GetAllProducts", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {

				found, err := l.svcCtx.Repo.Product().GetAllProducts(
					galaxyctx.GetStoreId(ctx),
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
						defaultImage := []*models.Asset{}
						defaultImage = append(defaultImage, &f.Asset)

						convertedImages := types.ConvertModelAssetToProtoAsset(defaultImage)
						if len(convertedImages) > 0 {
							prod.DefaultImage = convertedImages[0]
						}

						types.ConvertModelProductToProtoProduct(&f.Product, &[]models.Variant{
							f.Variant,
						}, &[]models.Price{
							f.Price,
						}, &prod)
						prods = append(prods, &prod)
					}
				}

				// the response struct
				item := &catalog.GetAllProductsResponse{
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

	res := &catalog.GetAllProductsResponse{}

	codec := &galaxycache.ByteCodec{}

	l.ctx = galaxyctx.SetStoreId(l.ctx, in.StoreId)
	l.ctx = galaxyctx.SetCurrentPage(l.ctx, in.CurrentPage)
	l.ctx = galaxyctx.SetPageSize(l.ctx, in.PageSize)
	l.ctx = galaxyctx.SetFilter(l.ctx, in.Filter)

	key := utils.StringToMD5(fmt.Sprintf("%d|%d|%d|%s", in.StoreId, in.CurrentPage, in.PageSize, in.Filter))
	if err := entryGetAllProductsLogic.galaxy.Get(l.ctx, key, codec); err != nil {
		return res, err
	}

	entryGetAllProductsLogic.galaxy.Remove(l.ctx, key)

	b, err := codec.MarshalBinary()
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(b, res)
	return res, err
}
