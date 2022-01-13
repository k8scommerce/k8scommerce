package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/internal/types"
	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetProductBySlugLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetProductBySlugLogic *galaxyGetProductBySlugLogicHelper

type GetProductBySlugLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetProductBySlugLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetProductBySlugLogic {
	return &GetProductBySlugLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetProductBySlugLogic) GetProductBySlug(in *catalog.GetProductBySlugRequest) (*catalog.GetProductBySlugResponse, error) {
	// caching goes logic here
	if entryGetProductBySlugLogic == nil {
		l.mu.Lock()
		entryGetProductBySlugLogic = &galaxyGetProductBySlugLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetProductBySlugLogic.once.Do(func() {
		// register the galaxy one time
		entryGetProductBySlugLogic.galaxy = gcache.RegisterGalaxyFunc("GetProductBySlug", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {

				v := strings.Split(key, "|")
				storeId, _ := strconv.ParseInt(v[0], 10, 64)
				slug := v[1]

				found, err := l.svcCtx.Repo.Product().GetProductBySlug(storeId, slug)
				if err != nil {
					logx.Infof("error: %s", err)
					return err
				}

				prod := catalog.Product{}
				if found != nil {
					types.ConvertModelProductToProtoProduct(&found.Product, &found.Variants, &found.Prices, &prod)
				}

				// the response struct
				item := &catalog.GetProductBySlugResponse{
					Product: &prod,
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
	entryGetProductBySlugLogic.galaxy.Get(l.ctx, key, codec)
	b, err := codec.MarshalBinary()
	if err != nil {
		return nil, err
	}
	res := &catalog.GetProductBySlugResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "",
	}
	err = json.Unmarshal(b, res)
	if err != nil {
		res.StatusCode = http.StatusExpectationFailed
		res.StatusMessage = err.Error()
	}
	return res, nil
}
