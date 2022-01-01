package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"store/internal/svc"
	"store/pb/store"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetAllStoresLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetAllStoresLogic *galaxyGetAllStoresLogicHelper

type GetAllStoresLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetAllStoresLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetAllStoresLogic {
	return &GetAllStoresLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetAllStoresLogic) GetAllStores(in *store.GetAllStoresRequest) (*store.GetAllStoresResponse, error) {

	// caching goes logic here
	if entryGetAllStoresLogic == nil {
		l.mu.Lock()
		entryGetAllStoresLogic = &galaxyGetAllStoresLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetAllStoresLogic.once.Do(func() {
		fmt.Println(`l.entryGetAllStoresLogic.Do`)

		// register the galaxy one time
		entryGetAllStoresLogic.galaxy = gcache.RegisterGalaxyFunc("GetAllStores", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				// todo: add your logic here and delete this line
				fmt.Printf("Looking up GetAllStores record by key: %s", key)

				// uncomment below to get the item from the adapter
				// found, err := l.ca.GetProductBySku(key)
				// if err != nil {
				//	logx.Infof("error: %s", err)
				//	return err
				// }

				// the response struct
				item := &store.GetAllStoresResponse{}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &store.GetAllStoresResponse{}

	codec := &galaxycache.ByteCodec{}
	if err := entryGetAllStoresLogic.galaxy.Get(l.ctx, in.Id, codec); err != nil {
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
