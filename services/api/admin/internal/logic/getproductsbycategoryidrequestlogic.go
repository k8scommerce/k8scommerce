package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"ecomm/services/api/admin/internal/svc"
	"ecomm/services/api/admin/internal/types"
	"ecomm/services/rpc/product/productclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductsByCategoryIdRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductsByCategoryIdRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductsByCategoryIdRequestLogic {
	return GetProductsByCategoryIdRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductsByCategoryIdRequestLogic) GetProductsByCategoryIdRequest(req types.GetProductsByCategoryIdRequest) (*types.GetProductsByCategoryIdResponse, error) {
	response, err := l.svcCtx.ProductRpc.GetProductsByCategoryId(l.ctx, &productclient.GetProductsByCategoryIdRequest{
		CategoryId:  req.CategoryId,
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
		SortOn:      req.SortOn,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("TotalRecords", response.TotalRecords)
	fmt.Println("TotalPages", response.TotalPages)

	// convert from one type to another
	// the structs are identical
	res := &types.GetProductsByCategoryIdResponse{}
	b, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, res)
	return res, err
}
