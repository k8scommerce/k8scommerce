package categories

import (
	"context"
	"encoding/json"
	"fmt"

	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
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

func (l *GetAllCategoriesLogic) GetAllCategories(req types.GetAllCategoriesRequest) (resp *types.GetAllCategoriesResponse, err error) {
	resp = &types.GetAllCategoriesResponse{}
	response, err := l.svcCtx.CatalogRpc.GetAllCategories(l.ctx, &catalogclient.GetAllCategoriesRequest{
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
		SortOn:      req.SortOn,
		StoreId:     l.ctx.Value(types.StoreKey).(int64),
	})
	fmt.Println(response)
	if err != nil {
		return nil, err
	}

	fmt.Println(response)

	// convert from one type to another
	// the structs are identical
	b, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &resp)
	return resp, err
}
