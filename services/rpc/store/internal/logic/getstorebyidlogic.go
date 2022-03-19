package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"k8scommerce/internal/convert"
	"k8scommerce/internal/galaxyctx"
	"k8scommerce/services/rpc/store/internal/svc"
	"k8scommerce/services/rpc/store/pb/store"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/zeromicro/go-zero/core/logx"
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

				found, err := l.svcCtx.Repo.Store().GetStoreById(galaxyctx.GetStoreId(ctx))
				if err != nil {
					logx.Infof("error: %s", err)
					return err
				}

				s := &store.Store{}
				if found != nil {
					convert.ModelStoreToProtoStore(found, s)
				}

				// the response struct
				item := &store.GetStoreByIdResponse{
					Store: s,
				}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &store.GetStoreByIdResponse{}

	l.ctx = galaxyctx.SetStoreId(l.ctx, in.Id)

	key := fmt.Sprintf("StoreID:%d", in.Id)
	codec := &galaxycache.ByteCodec{}
	if err := entryGetStoreByIdLogic.galaxy.Get(l.ctx, key, codec); err != nil {
		return res, nil
	}

	b, err := codec.MarshalBinary()
	if err != nil {
		return res, nil
	}

	err = json.Unmarshal(b, res)
	return res, err

}
