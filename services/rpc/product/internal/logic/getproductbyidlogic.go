package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/k8scommerce/k8scommerce/services/rpc/product/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/product/internal/types"
	"github.com/k8scommerce/k8scommerce/services/rpc/product/pb/product"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetProductByIdLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetProductByIdLogic *galaxyGetProductByIdLogicHelper

type GetProductByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetProductByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetProductByIdLogic {
	return &GetProductByIdLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetProductByIdLogic) GetProductById(in *product.GetProductByIdRequest) (*product.GetProductByIdResponse, error) {

	// caching goes logic here
	if entryGetProductByIdLogic == nil {
		l.mu.Lock()
		entryGetProductByIdLogic = &galaxyGetProductByIdLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetProductByIdLogic.once.Do(func() {
		// register the galaxy one time
		entryGetProductByIdLogic.galaxy = gcache.RegisterGalaxyFunc("GetProductById", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				fmt.Printf("Looking up GetProductById record by key: %s", key)

				id, _ := strconv.Atoi(key)
				found, err := l.svcCtx.Repo.Product().GetProductById(int64(id))
				if err != nil {
					logx.Infof("error: %s", err)
					return err
				}

				prod := product.Product{}
				if found != nil {
					types.ConvertModelProductToProtoProduct(&found.Product, &found.Variants, &found.Prices, &prod)
				}

				// the response struct
				item := &product.GetProductByIdResponse{
					Product: &prod,
				}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	codec := &galaxycache.ByteCodec{}
	entryGetProductByIdLogic.galaxy.Get(l.ctx, strconv.Itoa(int(in.Id)), codec)
	b, err := codec.MarshalBinary()
	if err != nil {
		return nil, err
	}
	res := &product.GetProductByIdResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "",
	}
	err = json.Unmarshal(b, res)
	if err != nil {
		res.StatusCode = http.StatusExpectationFailed
		res.StatusMessage = err.Error()
	}
	return res, nil
}
