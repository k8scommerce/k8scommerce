package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"k8scommerce/services/rpc/similarproducts/internal/svc"
	"k8scommerce/services/rpc/similarproducts/pb/similarproducts"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetSimilarProductsBySkuLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetSimilarProductsBySkuLogic *galaxyGetSimilarProductsBySkuLogicHelper

type GetSimilarProductsBySkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetSimilarProductsBySkuLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetSimilarProductsBySkuLogic {
	return &GetSimilarProductsBySkuLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetSimilarProductsBySkuLogic) GetSimilarProductsBySku(in *similarproducts.GetSimilarProductsBySkuRequest) (*similarproducts.GetSimilarProductsBySkuResponse, error) {

	// caching goes logic here
	if entryGetSimilarProductsBySkuLogic == nil {
		l.mu.Lock()
		entryGetSimilarProductsBySkuLogic = &galaxyGetSimilarProductsBySkuLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetSimilarProductsBySkuLogic.once.Do(func() {
		fmt.Println(`l.entryGetSimilarProductsBySkuLogic.Do`)

		// register the galaxy one time
		entryGetSimilarProductsBySkuLogic.galaxy = gcache.RegisterGalaxyFunc("GetSimilarProductsBySku", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				// todo: add your logic here and delete this line
				fmt.Printf("Looking up GetSimilarProductsBySku record by key: %s", key)

				// uncomment below to get the item from the adapter
				// found, err := l.ca.GetProductBySku(key)
				// if err != nil {
				//	logx.Infof("error: %s", err)
				//	return err
				// }

				// the response struct
				item := &similarproducts.GetSimilarProductsBySkuResponse{}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &similarproducts.GetSimilarProductsBySkuResponse{}

	codec := &galaxycache.ByteCodec{}
	if err := entryGetSimilarProductsBySkuLogic.galaxy.Get(l.ctx, in.Sku, codec); err != nil {
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
