package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"
	"net/http"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetCategoryBySlugLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetCategoryBySlugLogic *galaxyGetCategoryBySlugLogicHelper

type GetCategoryBySlugLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetCategoryBySlugLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetCategoryBySlugLogic {
	return &GetCategoryBySlugLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetCategoryBySlugLogic) GetCategoryBySlug(in *catalog.GetCategoryBySlugRequest) (*catalog.GetCategoryBySlugResponse, error) {

	// caching goes logic here
	if entryGetCategoryBySlugLogic == nil {
		l.mu.Lock()
		entryGetCategoryBySlugLogic = &galaxyGetCategoryBySlugLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetCategoryBySlugLogic.once.Do(func() {
		fmt.Println(`l.entryGetCategoryBySlugLogic.Do`)

		// register the galaxy one time
		entryGetCategoryBySlugLogic.galaxy = gcache.RegisterGalaxyFunc("GetCategoryBySlug", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				// todo: add your logic here and delete this line
				fmt.Printf("Looking up GetCategoryBySlug record by key: %s", key)

				// uncomment below to get the item from the adapter
				// found, err := l.ca.GetProductBySku(key)
				// if err != nil {
				//	logx.Infof("error: %s", err)
				//	return err
				// }

				// the response struct
				item := &catalog.GetCategoryBySlugResponse{}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &catalog.GetCategoryBySlugResponse{}

	codec := &galaxycache.ByteCodec{}
	if err := entryGetCategoryBySlugLogic.galaxy.Get(l.ctx, in.Slug, codec); err != nil {
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
