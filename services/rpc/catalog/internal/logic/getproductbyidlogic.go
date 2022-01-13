package logic

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/internal/types"
	"k8scommerce/services/rpc/catalog/pb/catalog"

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

func (l *GetProductByIdLogic) GetProductById(in *catalog.GetProductByIdRequest) (*catalog.GetProductByIdResponse, error) {

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
				id, _ := strconv.Atoi(key)
				found, err := l.svcCtx.Repo.Product().GetProductById(int64(id))
				if err != nil {
					logx.Infof("GetProductById error: %s", err)
					return err
				}

				prod := catalog.Product{}
				if found != nil {
					types.ConvertModelProductToProtoProduct(&found.Product, &found.Variants, &found.Prices, &prod)
				}

				// the response struct
				item := &catalog.GetProductByIdResponse{
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
	res := &catalog.GetProductByIdResponse{
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
