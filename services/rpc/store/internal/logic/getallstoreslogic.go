package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"k8scommerce/services/rpc/store/internal/svc"
	"k8scommerce/services/rpc/store/pb/store"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/zeromicro/go-zero/core/logx"
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
				// found, err := l.svcCtx.Repo.Store().GetStoreById(galaxyctx.GetStoreId(ctx))
				// if err != nil {
				// 	logx.Infof("error: %s", err)
				// 	return err
				// }

				// s := &store.Store{}
				// if found != nil {
				//  convert.ModelStoreToProtoStore(found, s)
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
	if err := entryGetAllStoresLogic.galaxy.Get(l.ctx, "all-stores", codec); err != nil {
		return res, nil
	}

	b, err := codec.MarshalBinary()
	if err != nil {
		return res, nil
	}

	err = json.Unmarshal(b, res)
	return res, err

}
