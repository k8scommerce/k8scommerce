package logic

import (
	"context"
	"encoding/json"
	"fmt"
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
				storeId, _ := strconv.ParseInt(v[0], 10, 64)
				currentPage, _ := strconv.ParseInt(v[1], 10, 64)
				pageSize, _ := strconv.ParseInt(v[2], 10, 64)
				sortOn := ""
				if len(v) > 3 {
					sortOn = v[3]
				}
				found, err := l.svcCtx.Repo.Category().GetAllCategories(storeId, currentPage, pageSize, sortOn)
				if err != nil {
					logx.Infof("error: %s", err)
					return err
				}

				cats := []*catalog.Category{}

				var totalRecords int64 = 0
				var totalPages int64 = 0

				if found != nil {
					totalRecords = found.PagingStats.TotalRecords
					totalPages = found.PagingStats.TotalPages

					for _, f := range found.Categories {
						cat := catalog.Category{}
						types.ConvertModelCategoryToProtoCategory(&f, &cat)
						cats = append(cats, &cat)
					}
				}

				// the response struct
				item := &catalog.GetAllCategoriesResponse{
					Categories:   cats,
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
	key := fmt.Sprintf("%d|%d|%d|%s", in.StoreId, in.CurrentPage, in.PageSize, in.SortOn)
	if err := entryGetAllCategoriesLogic.galaxy.Get(l.ctx, key, codec); err != nil {
		// res.StatusCode = http.StatusNoContent
		// res.StatusMessage = "ERROR 2: " + err.Error()
		return res, err
	}

	b, err := codec.MarshalBinary()
	if err != nil {
		// res.StatusCode = http.StatusInternalServerError
		// res.StatusMessage = "ERROR 2: " + err.Error()
		return res, err
	}

	err = json.Unmarshal(b, res)

	// remove for testing
	entryGetAllCategoriesLogic.galaxy.Remove(l.ctx, key)

	return res, err

}
