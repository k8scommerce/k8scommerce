package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"warehouse/internal/svc"
	"warehouse/pb/warehouse"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetAllWarehousesByStoreIdLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetAllWarehousesByStoreIdLogic *galaxyGetAllWarehousesByStoreIdLogicHelper

type GetAllWarehousesByStoreIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetAllWarehousesByStoreIdLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetAllWarehousesByStoreIdLogic {
	return &GetAllWarehousesByStoreIdLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetAllWarehousesByStoreIdLogic) GetAllWarehousesByStoreId(in *warehouse.GetAllWarehousesByStoreIdRequest) (*warehouse.GetAllWarehousesByStoreIdResponse, error) {

	// caching goes logic here
	if entryGetAllWarehousesByStoreIdLogic == nil {
		l.mu.Lock()
		entryGetAllWarehousesByStoreIdLogic = &galaxyGetAllWarehousesByStoreIdLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetAllWarehousesByStoreIdLogic.once.Do(func() {
		fmt.Println(`l.entryGetAllWarehousesByStoreIdLogic.Do`)

		// register the galaxy one time
		entryGetAllWarehousesByStoreIdLogic.galaxy = gcache.RegisterGalaxyFunc("GetAllWarehousesByStoreId", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				// todo: add your logic here and delete this line
				fmt.Printf("Looking up GetAllWarehousesByStoreId record by key: %s", key)

				// uncomment below to get the item from the adapter
				// found, err := l.ca.GetProductBySku(key)
				// if err != nil {
				//	logx.Infof("error: %s", err)
				//	return err
				// }

				// the response struct
				item := &warehouse.GetAllWarehousesByStoreIdResponse{}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &warehouse.GetAllWarehousesByStoreIdResponse{}

	codec := &galaxycache.ByteCodec{}
	if err := entryGetAllWarehousesByStoreIdLogic.galaxy.Get(l.ctx, in.Id, codec); err != nil {
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
