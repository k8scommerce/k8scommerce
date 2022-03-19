package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"k8scommerce/internal/convert"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"
	"strconv"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/zeromicro/go-zero/core/logx"
)

type galaxyGetCategoryByIdLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetCategoryByIdLogic *galaxyGetCategoryByIdLogicHelper

type GetCategoryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetCategoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetCategoryByIdLogic {
	return &GetCategoryByIdLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetCategoryByIdLogic) GetCategoryById(in *catalog.GetCategoryByIdRequest) (*catalog.GetCategoryByIdResponse, error) {

	// caching goes logic here
	if entryGetCategoryByIdLogic == nil {
		l.mu.Lock()
		entryGetCategoryByIdLogic = &galaxyGetCategoryByIdLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetCategoryByIdLogic.once.Do(func() {
		// register the galaxy one time
		entryGetCategoryByIdLogic.galaxy = gcache.RegisterGalaxyFunc("GetCategoryById", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {
				fmt.Printf("Looking up GetCategoryById record by key: %s", key)

				categoryId, _ := strconv.ParseInt(key, 10, 64)
				found, err := l.svcCtx.Repo.Category().GetCategoryById(categoryId)
				if err != nil {
					logx.Infof("error: %s", err)
					return err
				}

				cat := catalog.Category{}
				if found != nil {
					convert.ModelCategoryToProtoCategory(found, &cat)
				}

				// the response struct
				item := &catalog.GetCategoryByIdResponse{
					Category: &cat,
				}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	codec := &galaxycache.ByteCodec{}
	entryGetCategoryByIdLogic.galaxy.Get(l.ctx, strconv.Itoa(int(in.Id)), codec)
	b, err := codec.MarshalBinary()
	if err != nil {
		return nil, err
	}
	res := &catalog.GetCategoryByIdResponse{}
	err = json.Unmarshal(b, res)
	return res, err
}
