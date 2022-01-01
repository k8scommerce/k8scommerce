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

type galaxyGetStoreByIdLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetStoreByIdLogic *galaxyGetStoreByIdLogicHelper

type GetStoreByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetStoreByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetStoreByIdLogic {
	return &GetStoreByIdLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetStoreByIdLogic) GetStoreById(in *store.GetStoreByIdRequest) (*store.GetStoreByIdResponse, error) {

	// caching goes logic here
	if entryGetStoreByIdLogic == nil {
		l.mu.Lock()
		entryGetStoreByIdLogic = &galaxyGetStoreByIdLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetStoreByIdLogic.once.Do(func() {
		fmt.Println(`l.entryGetStoreByIdLogic.Do`)

		// register the galaxy one time
		entryGetStoreByIdLogic.galaxy = gcache.RegisterGalaxyFunc("GetStoreById", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				// todo: add your logic here and delete this line
				fmt.Printf("Looking up GetStoreById record by key: %s", key)

				// uncomment below to get the item from the adapter
				// found, err := l.ca.GetProductBySku(key)
				// if err != nil {
				//	logx.Infof("error: %s", err)
				//	return err
				// }

				// the response struct
				item := &store.GetStoreByIdResponse{}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &store.GetStoreByIdResponse{}

	codec := &galaxycache.ByteCodec{}
	if err := entryGetStoreByIdLogic.galaxy.Get(l.ctx, in.Id, codec); err != nil {
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
