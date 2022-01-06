package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/internal/types"
	"k8scommerce/services/rpc/catalog/pb/catalog"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
)

const (
	AllCatgoriesKey = "all-categories"
)

type galaxyGetAllCategoriesLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetAllCategoriesLogic *galaxyGetAllCategoriesLogicHelper

type GetAllCategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetAllCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetAllCategoriesLogic {
	return &GetAllCategoriesLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

//  categories
func (l *GetAllCategoriesLogic) GetAllCategories(in *catalog.GetAllCategoriesRequest) (*catalog.GetAllCategoriesResponse, error) {

	// caching goes logic here
	if entryGetAllCategoriesLogic == nil {
		l.mu.Lock()
		entryGetAllCategoriesLogic = &galaxyGetAllCategoriesLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetAllCategoriesLogic.once.Do(func() {
		fmt.Println(`l.entryGetAllCategoriesLogic.Do`)

		// register the galaxy one time
		entryGetAllCategoriesLogic.galaxy = gcache.RegisterGalaxyFunc("GetAllCategories", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				// split the key and set the variables
				v := strings.Split(key, "|")
				currentPage, _ := strconv.ParseInt(v[1], 10, 64)
				pageSize, _ := strconv.ParseInt(v[2], 10, 64)
				sortOn := ""
				if len(v) > 3 {
					sortOn = v[3]
				}
				found, err := l.svcCtx.Repo.Category().GetAllCategories(currentPage, pageSize, sortOn)
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

	res := &catalog.GetAllCategoriesResponse{}

	codec := &galaxycache.ByteCodec{}
	if err := entryGetAllCategoriesLogic.galaxy.Get(l.ctx, AllCatgoriesKey, codec); err != nil {
		res.StatusCode = http.StatusNoContent
		res.StatusMessage = err.Error()
		return res, nil
	}

	b, err := codec.MarshalBinary()
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.StatusMessage = err.Error()
		return res, nil
	}

	err = json.Unmarshal(b, res)
	return res, err

}
