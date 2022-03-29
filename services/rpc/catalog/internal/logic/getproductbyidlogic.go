package logic

import (
	"context"
	"time"

	"k8scommerce/internal/convert"
	"k8scommerce/internal/gcache"
	"k8scommerce/internal/groupctx"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

const Group_GetProductById = "GetProductById"

var Group_GetProductByIdKey = func(productId int64) string {
	return gcache.ToKey(Group_GetProductById, productId)
}

type GetProductByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductByIdLogic {
	return &GetProductByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductByIdLogic) GetProductById(in *catalog.GetProductByIdRequest) (*catalog.GetProductByIdResponse, error) {
	l.ctx = groupctx.SetProductId(l.ctx, in.Id)
	res := &catalog.GetProductByIdResponse{}
	err := l.cache().Get(l.ctx, Group_GetProductByIdKey(in.Id), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetProductByIdLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetAllCategories, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.Product().GetProductById(
				groupctx.GetProductId(ctx),
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

			// Set the groupcache to expire after 24 hours
			if err := dest.SetProto(&catalog.GetProductByIdResponse{
				Product: &prod,
			}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
