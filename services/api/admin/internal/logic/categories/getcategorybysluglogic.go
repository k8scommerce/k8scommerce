package categories

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/types"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/catalogclient"

	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryBySlugLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryBySlugLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCategoryBySlugLogic {
	return GetCategoryBySlugLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryBySlugLogic) GetCategoryBySlug(req types.GetCategoryBySlugRequest) (resp *types.Category, err error) {
	resp = &types.Category{}

	response, err := l.svcCtx.CatalogRpc.GetCategoryBySlug(l.ctx, &catalogclient.GetCategoryBySlugRequest{
		Slug:    req.Slug,
		StoreId: l.ctx.Value(types.StoreKey).(int64),
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	utils.TransformObj(response.Category, &resp)

	return resp, err
}
