package logic

import (
	"context"
	"encoding/json"
	"net/http"

	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
	"k8scommerce/services/rpc/catalog/catalogclient"

	"github.com/tal-tech/go-zero/core/logx"
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
		StoreId:     req.StoreId,
	})
	if err != nil {
		resp.ResponseStatus = httpResponse(http.StatusBadRequest, err.Error())
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	b, err := json.Marshal(response)
	if err != nil {
		resp.ResponseStatus = httpResponse(http.StatusBadRequest, err.Error())
		return nil, err
	}
	err = json.Unmarshal(b, &resp)
	resp.ResponseStatus = httpResponse(http.StatusOK)
	return resp, err
}
