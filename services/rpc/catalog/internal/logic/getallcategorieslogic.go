package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"k8scommerce/internal/galaxyctx"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/internal/types"
	"k8scommerce/services/rpc/catalog/pb/catalog"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/zeromicro/go-zero/core/logx"
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

		// register the galaxy one time
		entryGetAllCategoriesLogic.galaxy = gcache.RegisterGalaxyFunc("GetAllCategories", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {

				found, err := l.svcCtx.Repo.Category().GetAllCategories(
					galaxyctx.GetStoreId(ctx),
				)
				if err != nil {
					logx.Infof("error: %s", err)
					return err
				}

				cats := []*catalog.Category{}

				if found != nil {
					for _, f := range found.Categories {
						cat := catalog.Category{}
						types.ConvertModelCategoryToProtoCategory(&f, &cat)
						cats = append(cats, &cat)
					}
				}

				// the response struct
				item := &catalog.GetAllCategoriesResponse{
					Categories: cats,
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

	l.ctx = galaxyctx.SetStoreId(l.ctx, in.StoreId)

	key := fmt.Sprintf("categories-getall-%d", in.StoreId)
	if err := entryGetAllCategoriesLogic.galaxy.Get(l.ctx, key, codec); err != nil {
		return res, err
	}

	b, err := codec.MarshalBinary()
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(b, res)

	// remove for testing
	entryGetAllCategoriesLogic.galaxy.Remove(l.ctx, key)

	return res, err

}
