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
		// register the galaxy one time
		entryGetCategoryBySlugLogic.galaxy = gcache.RegisterGalaxyFunc("GetCategoryBySlug", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				fmt.Printf("Looking up GetCategoryBySlug record by key: %s", key)

				v := strings.Split(key, "|")
				storeId, _ := strconv.ParseInt(v[1], 10, 64)
				slug := v[1]

				found, err := l.svcCtx.Repo.Category().GetCategoryBySlug(storeId, slug)
				if err != nil {
					logx.Infof("error: %s", err)
					return err
				}

				cat := catalog.Category{}
				if found != nil {
					types.ConvertModelCategoryToProtoCategory(found, &cat)
				}

				// the response struct
				item := &catalog.GetCategoryBySlugResponse{
					Category: &cat,
				}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	codec := &galaxycache.ByteCodec{}
	key := fmt.Sprintf("%d|%s", in.StoreId, in.Slug)
	entryGetCategoryBySlugLogic.galaxy.Get(l.ctx, key, codec)
	b, err := codec.MarshalBinary()
	if err != nil {
		return nil, err
	}
	res := &catalog.GetCategoryBySlugResponse{}
	err = json.Unmarshal(b, res)
	return res, err
}
