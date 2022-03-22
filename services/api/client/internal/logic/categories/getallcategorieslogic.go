package categories

import (
	"context"

	"k8scommerce/internal/utils"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
	"k8scommerce/services/rpc/catalog/catalogclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllCategoriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllCategoriesLogic {
	return GetAllCategoriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllCategoriesLogic) GetAllCategories() (resp *types.GetAllCategoriesResponse, err error) {
	response, err := l.svcCtx.CatalogRpc.GetAllCategories(l.ctx, &catalogclient.GetAllCategoriesRequest{
		StoreId: l.ctx.Value(types.StoreKey).(int64),
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	utils.TransformObj(response, &resp)
	return resp, err
}
