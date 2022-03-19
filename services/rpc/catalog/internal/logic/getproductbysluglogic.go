package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"k8scommerce/internal/convert"
	"k8scommerce/internal/galaxyctx"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/zeromicro/go-zero/core/logx"
)

type galaxyGetProductBySlugLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetProductBySlugLogic *galaxyGetProductBySlugLogicHelper

type GetProductBySlugLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetProductBySlugLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetProductBySlugLogic {
	return &GetProductBySlugLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetProductBySlugLogic) GetProductBySlug(in *catalog.GetProductBySlugRequest) (*catalog.GetProductBySlugResponse, error) {
	// caching goes logic here
	if entryGetProductBySlugLogic == nil {
		l.mu.Lock()
		entryGetProductBySlugLogic = &galaxyGetProductBySlugLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetProductBySlugLogic.once.Do(func() {
		// register the galaxy one time
		entryGetProductBySlugLogic.galaxy = gcache.RegisterGalaxyFunc("GetProductBySlug", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {

				found, err := l.svcCtx.Repo.Product().GetProductBySlug(
					galaxyctx.GetStoreId(ctx),
					galaxyctx.GetSlug(ctx),
				)
				if err != nil {
					logx.Infof("error: %s", err)
					return err
				}

				prod := catalog.Product{}
				if found != nil {
					convert.ModelProductToProtoProduct(&found.Product, &found.Variants, &found.Prices, &prod)

					for _, pair := range found.Categories {
						prod.Categories = append(prod.Categories, &catalog.CategoryPair{
							Slug: pair.Slug,
							Name: pair.Name,
						})
					}
				}

				// get the images
				images, err := l.svcCtx.Repo.Asset().GetAssetByProductIDKind(prod.Id, int(catalog.AssetKind_image))
				if err != nil {
					logx.Infof("error: %s", err)
					return err
				}
				if images != nil {
					prod.Images = convert.ModelAssetToProtoAsset(images)
				}

				// the response struct
				item := &catalog.GetProductBySlugResponse{
					Product: &prod,
				}

				out, err := json.Marshal(item)
				if err != nil {
					logx.Infof("error: %s", err)
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	codec := &galaxycache.ByteCodec{}

	l.ctx = galaxyctx.SetStoreId(l.ctx, in.StoreId)
	l.ctx = galaxyctx.SetSlug(l.ctx, in.Slug)

	key := fmt.Sprintf("%d|%s", in.StoreId, in.Slug)
	entryGetProductBySlugLogic.galaxy.Get(l.ctx, key, codec)
	b, err := codec.MarshalBinary()
	if err != nil {
		logx.Infof("error: %s", err)
		return nil, err
	}

	entryGetProductBySlugLogic.galaxy.Remove(l.ctx, key)

	res := &catalog.GetProductBySlugResponse{
		Product: &catalog.Product{},
	}
	err = json.Unmarshal(b, res)
	return res, err
}
