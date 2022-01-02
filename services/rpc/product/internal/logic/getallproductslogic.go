package logic

import (
	"context"

	"encoding/json"
	"fmt"
	"sync"

	"github.com/k8scommerce/k8scommerce/services/rpc/product/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/product/pb/product"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetAllProductsLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetAllProductsLogic map[string]*galaxyGetAllProductsLogicHelper

func init() {
	entryGetAllProductsLogic = make(map[string]*galaxyGetAllProductsLogicHelper)
}

type GetAllProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetAllProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetAllProductsLogic {
	return &GetAllProductsLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetAllProductsLogic) GetAllProducts(in *product.GetAllProductsRequest) (*product.GetAllProductsResponse, error) {

	// caching goes logic here
	if _, ok := entryGetAllProductsLogic["GetAllProducts"]; !ok {
		l.mu.Lock()
		entryGetAllProductsLogic["GetAllProducts"] = &galaxyGetAllProductsLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetAllProductsLogic["GetAllProducts"].once.Do(func() {
		fmt.Println(`l.entry["GetAllProducts"].Do`)

		// register the galaxy one time
		entryGetAllProductsLogic["GetAllProducts"].galaxy = gcache.RegisterGalaxyFunc("GetAllProducts", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				// todo: add your logic here and delete this line
				fmt.Printf("Looking up GetAllProducts record by key: %s", key)

				// uncomment below to get the item from the adapter
				// found, err := l.ca.GetProductBySku(key)
				// if err != nil {
				//	logx.Infof("error: %s", err)
				//	return err
				// }

				// the response struct
				item := &product.GetAllProductsResponse{}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	codec := &galaxycache.ByteCodec{}
	// entryGetAllProductsLogic["GetAllProducts"].galaxy.Get(l.ctx, strconv.Itoa(int(in.Id)), codec)
	b, err := codec.MarshalBinary()
	if err != nil {
		return nil, err
	}
	res := &product.GetAllProductsResponse{}
	err = json.Unmarshal(b, res)
	return res, err

}
