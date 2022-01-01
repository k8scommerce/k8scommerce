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

type galaxyCreateStoreLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryCreateStoreLogic *galaxyCreateStoreLogicHelper

type CreateStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewCreateStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *CreateStoreLogic {
	return &CreateStoreLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *CreateStoreLogic) CreateStore(in *store.CreateStoreRequest) (*store.CreateStoreResponse, error) {

	// caching goes logic here
	if entryCreateStoreLogic == nil {
		l.mu.Lock()
		entryCreateStoreLogic = &galaxyCreateStoreLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryCreateStoreLogic.once.Do(func() {
		fmt.Println(`l.entryCreateStoreLogic.Do`)

		// register the galaxy one time
		entryCreateStoreLogic.galaxy = gcache.RegisterGalaxyFunc("CreateStore", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				// todo: add your logic here and delete this line
				fmt.Printf("Looking up CreateStore record by key: %s", key)

				// uncomment below to get the item from the adapter
				// found, err := l.ca.GetProductBySku(key)
				// if err != nil {
				//	logx.Infof("error: %s", err)
				//	return err
				// }

				// the response struct
				item := &store.CreateStoreResponse{}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &store.CreateStoreResponse{}

	codec := &galaxycache.ByteCodec{}
	if err := entryCreateStoreLogic.galaxy.Get(l.ctx, in.Id, codec); err != nil {
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
