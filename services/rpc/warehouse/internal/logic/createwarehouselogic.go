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

type galaxyCreateWarehouseLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryCreateWarehouseLogic *galaxyCreateWarehouseLogicHelper

type CreateWarehouseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewCreateWarehouseLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *CreateWarehouseLogic {
	return &CreateWarehouseLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *CreateWarehouseLogic) CreateWarehouse(in *warehouse.CreateWarehouseRequest) (*warehouse.CreateWarehouseResponse, error) {

	// caching goes logic here
	if entryCreateWarehouseLogic == nil {
		l.mu.Lock()
		entryCreateWarehouseLogic = &galaxyCreateWarehouseLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryCreateWarehouseLogic.once.Do(func() {
		fmt.Println(`l.entryCreateWarehouseLogic.Do`)

		// register the galaxy one time
		entryCreateWarehouseLogic.galaxy = gcache.RegisterGalaxyFunc("CreateWarehouse", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				// todo: add your logic here and delete this line
				fmt.Printf("Looking up CreateWarehouse record by key: %s", key)

				// uncomment below to get the item from the adapter
				// found, err := l.ca.GetProductBySku(key)
				// if err != nil {
				//	logx.Infof("error: %s", err)
				//	return err
				// }

				// the response struct
				item := &warehouse.CreateWarehouseResponse{}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &warehouse.CreateWarehouseResponse{}

	codec := &galaxycache.ByteCodec{}
	if err := entryCreateWarehouseLogic.galaxy.Get(l.ctx, in.Id, codec); err != nil {
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
