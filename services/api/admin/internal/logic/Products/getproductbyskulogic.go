package Products

import (
	"context"

	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProductBySkuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductBySkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetProductBySkuLogic {
	return GetProductBySkuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductBySkuLogic) GetProductBySku(req types.GetProductBySkuRequest) (resp *types.GetProductResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
