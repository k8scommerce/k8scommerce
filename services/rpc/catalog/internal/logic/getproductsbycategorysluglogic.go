package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/internal/types"
	"k8scommerce/services/rpc/catalog/pb/catalog"
	"strconv"
	"strings"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
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
		// fmt.Println(`l.entry.Do`)

		// register the galaxy one time
		entryGetProductsByCategorySlugLogic.galaxy = gcache.RegisterGalaxyFunc("GetProductsByCategorySlug", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {

				// split the key and set the variables

				v := strings.Split(key, "|")
				storeId, _ := strconv.ParseInt(v[0], 10, 64)
				categorySlug := v[1]
				currentPage, _ := strconv.ParseInt(v[2], 10, 64)
				pageSize, _ := strconv.ParseInt(v[3], 10, 64)
				sortOn := ""
				if len(v) > 4 {
					sortOn = v[4]
				}
				found, err := l.svcCtx.Repo.Product().GetProductsByCategorySlug(storeId, categorySlug, currentPage, pageSize, sortOn)
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

						types.ConvertModelProductToProtoProduct(&f.Product, &[]models.Variant{
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

	key := fmt.Sprintf("%d|%s|%d|%d|%s", in.StoreId, in.CategorySlug, in.CurrentPage, in.PageSize, in.SortOn)
	entryGetProductsByCategorySlugLogic.galaxy.Get(l.ctx, key, codec)
	b, err := codec.MarshalBinary()
	if err != nil {
		return nil, err
	}
	res := &catalog.GetProductsByCategorySlugResponse{}
	err = json.Unmarshal(b, res)

	return res, err
}
