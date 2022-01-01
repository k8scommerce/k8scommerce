package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/k8s-commerce/k8s-commerce/services/rpc/othersbought/internal/svc"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/othersbought/pb/othersbought"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetOthersBoughtBySkuLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetOthersBoughtBySkuLogic *galaxyGetOthersBoughtBySkuLogicHelper

type GetOthersBoughtBySkuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetOthersBoughtBySkuLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetOthersBoughtBySkuLogic {
	return &GetOthersBoughtBySkuLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetOthersBoughtBySkuLogic) GetOthersBoughtBySku(in *othersbought.GetOthersBoughtBySkuRequest) (*othersbought.GetOthersBoughtBySkuResponse, error) {

	// caching goes logic here
	if entryGetOthersBoughtBySkuLogic == nil {
		l.mu.Lock()
		entryGetOthersBoughtBySkuLogic = &galaxyGetOthersBoughtBySkuLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetOthersBoughtBySkuLogic.once.Do(func() {
		fmt.Println(`l.entryGetOthersBoughtBySkuLogic.Do`)

		// register the galaxy one time
		entryGetOthersBoughtBySkuLogic.galaxy = gcache.RegisterGalaxyFunc("GetOthersBoughtBySku", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				// todo: add your logic here and delete this line
				fmt.Printf("Looking up GetOthersBoughtBySku record by key: %s", key)

				// uncomment below to get the item from the adapter
				// found, err := l.ca.GetProductBySku(key)
				// if err != nil {
				//	logx.Infof("error: %s", err)
				//	return err
				// }

				// the response struct
				item := &othersbought.GetOthersBoughtBySkuResponse{}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &othersbought.GetOthersBoughtBySkuResponse{}

	codec := &galaxycache.ByteCodec{}
	if err := entryGetOthersBoughtBySkuLogic.galaxy.Get(l.ctx, in.Sku, codec); err != nil {
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
