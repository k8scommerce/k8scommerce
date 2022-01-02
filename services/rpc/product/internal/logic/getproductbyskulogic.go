package logic

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/k8scommerce/k8scommerce/services/rpc/product/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/product/internal/types"
	"github.com/k8scommerce/k8scommerce/services/rpc/product/pb/product"

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

func (l *GetProductBySkuLogic) GetProductBySku(in *product.GetProductBySkuRequest) (*product.GetProductBySkuResponse, error) {

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
				// fmt.Printf("Looking up GetProductBySku record by key: %s", key)
				found, err := l.svcCtx.Repo.Product().GetProductBySku(key)
				if err != nil {
					logx.Infof("error: %s", err)
					return err
				}

				prod := product.Product{}
				if found != nil {
					types.ConvertModelProductToProtoProduct(&found.Product, &found.Variants, &found.Prices, &prod)
				}

				// the response struct
				item := &product.GetProductBySkuResponse{
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
	entryGetProductBySkuLogic.galaxy.Get(l.ctx, in.Sku, codec)
	b, err := codec.MarshalBinary()
	if err != nil {
		return nil, err
	}
	res := &product.GetProductBySkuResponse{
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
