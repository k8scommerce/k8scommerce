package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"k8scommerce/internal/convert"
	"k8scommerce/internal/galaxyctx"
	"k8scommerce/services/rpc/store/internal/svc"
	"k8scommerce/services/rpc/store/pb/store"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/zeromicro/go-zero/core/logx"
)

type galaxyGetStoreSettingsLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetStoreSettingsLogic *galaxyGetStoreSettingsLogicHelper

type GetStoreSettingsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetStoreSettingsLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetStoreSettingsLogic {
	return &GetStoreSettingsLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetStoreSettingsLogic) GetStoreSettings(in *store.GetStoreSettingRequest) (*store.GetStoreSettingResponse, error) {

	// caching goes logic here
	if entryGetStoreSettingsLogic == nil {
		l.mu.Lock()
		entryGetStoreSettingsLogic = &galaxyGetStoreSettingsLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetStoreSettingsLogic.once.Do(func() {
		fmt.Println(`l.entryGetStoreSettingsLogic.Do`)

		// register the galaxy one time
		entryGetStoreSettingsLogic.galaxy = gcache.RegisterGalaxyFunc("GetStoreSettings", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {

				// uncomment below to get the item from the adapter
				found, err := l.svcCtx.Repo.StoreSetting().GetStoreSettingById(galaxyctx.GetStoreId(ctx))
				if err != nil {
					logx.Infof("error: %s", err)
					return err
				}

				setting := &store.StoreSetting{}
				if found != nil {
					convert.ModelStoreSettingToProtoStoreSetting(found, setting)
				}

				// the response struct
				item := &store.GetStoreSettingResponse{
					Setting: setting,
				}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &store.GetStoreSettingResponse{}

	l.ctx = galaxyctx.SetStoreId(l.ctx, in.StoreId)
	key := fmt.Sprintf("StoreID:%d", in.StoreId)
	codec := &galaxycache.ByteCodec{}
	if err := entryGetStoreSettingsLogic.galaxy.Get(l.ctx, key, codec); err != nil {
		return res, nil
	}

	b, err := codec.MarshalBinary()
	if err != nil {
		return res, nil
	}

	err = json.Unmarshal(b, res)
	return res, err

}
