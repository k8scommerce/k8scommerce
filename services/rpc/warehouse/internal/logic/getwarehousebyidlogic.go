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

type galaxyGetWarehouseByIdLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetWarehouseByIdLogic *galaxyGetWarehouseByIdLogicHelper

type GetWarehouseByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetWarehouseByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetWarehouseByIdLogic {
	return &GetWarehouseByIdLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetWarehouseByIdLogic) GetWarehouseById(in *warehouse.GetWarehouseByIdRequest) (*warehouse.GetWarehouseByIdResponse, error) {

	// caching goes logic here
	if entryGetWarehouseByIdLogic == nil {
		l.mu.Lock()
		entryGetWarehouseByIdLogic = &galaxyGetWarehouseByIdLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetWarehouseByIdLogic.once.Do(func() {
		fmt.Println(`l.entryGetWarehouseByIdLogic.Do`)

		// register the galaxy one time
		entryGetWarehouseByIdLogic.galaxy = gcache.RegisterGalaxyFunc("GetWarehouseById", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				// todo: add your logic here and delete this line
				fmt.Printf("Looking up GetWarehouseById record by key: %s", key)

				// uncomment below to get the item from the adapter
				// found, err := l.ca.GetProductBySku(key)
				// if err != nil {
				//	logx.Infof("error: %s", err)
				//	return err
				// }

				// the response struct
				item := &warehouse.GetWarehouseByIdResponse{}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &warehouse.GetWarehouseByIdResponse{}

	codec := &galaxycache.ByteCodec{}
	if err := entryGetWarehouseByIdLogic.galaxy.Get(l.ctx, in.Id, codec); err != nil {
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
