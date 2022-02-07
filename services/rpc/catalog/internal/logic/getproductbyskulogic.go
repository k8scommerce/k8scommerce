package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"k8scommerce/internal/galaxyctx"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/internal/types"
	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetProductBySkuLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetProductBySkuLogic *galaxyGetProductBySkuLogicHelper

type GetProductBySkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetProductBySkuLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetProductBySkuLogic {
	return &GetProductBySkuLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetProductBySkuLogic) GetProductBySku(in *catalog.GetProductBySkuRequest) (*catalog.GetProductBySkuResponse, error) {
	// caching goes logic here
	if entryGetProductBySkuLogic == nil {
		l.mu.Lock()
		entryGetProductBySkuLogic = &galaxyGetProductBySkuLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetProductBySkuLogic.once.Do(func() {
		// register the galaxy one time
		entryGetProductBySkuLogic.galaxy = gcache.RegisterGalaxyFunc("GetProductBySku", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {

				found, err := l.svcCtx.Repo.Product().GetProductBySku(
					galaxyctx.GetStoreId(ctx),
					galaxyctx.GetSku(ctx),
				)
				if err != nil {
					logx.Infof("error: %s", err)
					return err
				}

				prod := catalog.Product{}
				if found != nil {
					types.ConvertModelProductToProtoProduct(&found.Product, &found.Variants, &found.Prices, &prod)
				}

				// the response struct
				item := &catalog.GetProductBySkuResponse{
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

	l.ctx = galaxyctx.SetStoreId(l.ctx, in.StoreId)
	l.ctx = galaxyctx.SetSku(l.ctx, in.Sku)

	key := fmt.Sprintf("%d|%s", in.StoreId, in.Sku)
	entryGetProductBySkuLogic.galaxy.Get(l.ctx, key, codec)
	b, err := codec.MarshalBinary()
	if err != nil {
		return nil, err
	}
	res := &catalog.GetProductBySkuResponse{}
	err = json.Unmarshal(b, res)
	return res, err
}
