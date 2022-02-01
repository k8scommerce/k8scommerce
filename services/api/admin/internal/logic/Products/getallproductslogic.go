package Products

import (
	"context"

	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllProductsLogic {
	return GetAllProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllProductsLogic) GetAllProducts(req types.GetAllProductsRequest) (resp *types.GetAllProductsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
