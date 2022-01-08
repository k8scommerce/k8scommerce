package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/internal/types"
	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetProductsByCategoryIdLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetProductsByCategoryIdLogic *galaxyGetProductsByCategoryIdLogicHelper

type GetProductsByCategoryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetProductsByCategoryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetProductsByCategoryIdLogic {
	return &GetProductsByCategoryIdLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetProductsByCategoryIdLogic) GetProductsByCategoryId(in *catalog.GetProductsByCategoryIdRequest) (*catalog.GetProductsByCategoryIdResponse, error) {

	// caching goes logic here
	if entryGetProductsByCategoryIdLogic == nil {
		l.mu.Lock()
		entryGetProductsByCategoryIdLogic = &galaxyGetProductsByCategoryIdLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetProductsByCategoryIdLogic.once.Do(func() {
		// fmt.Println(`l.entry.Do`)

		// register the galaxy one time
		entryGetProductsByCategoryIdLogic.galaxy = gcache.RegisterGalaxyFunc("GetProductsByCategoryId", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {

				// split the key and set the variables

				v := strings.Split(key, "|")
				storeId, _ := strconv.ParseInt(v[0], 10, 64)
				categoryId, _ := strconv.ParseInt(v[1], 10, 64)
				currentPage, _ := strconv.ParseInt(v[2], 10, 64)
				pageSize, _ := strconv.ParseInt(v[3], 10, 64)
				sortOn := ""
				if len(v) > 4 {
					sortOn = v[4]
				}
				found, err := l.svcCtx.Repo.Product().GetProductsByCategoryId(storeId, categoryId, currentPage, pageSize, sortOn)
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
				item := &catalog.GetProductsByCategoryIdResponse{
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

	key := fmt.Sprintf("%d|%d|%d|%d|%s", in.StoreId, in.CurrentPage, in.PageSize, in.SortOn)
	// fmt.Println(key, in.CategoryId, in.CurrentPage, in.PageSize, in.SortOn)
	entryGetProductsByCategoryIdLogic.galaxy.Get(l.ctx, key, codec)
	b, err := codec.MarshalBinary()
	if err != nil {
		return nil, err
	}
	res := &catalog.GetProductsByCategoryIdResponse{}
	err = json.Unmarshal(b, res)

	// remove it for right now
	// entryGetProductsByCategoryIdLogic.galaxy.Remove(l.ctx, key)

	return res, err

}
