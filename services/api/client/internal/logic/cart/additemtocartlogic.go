package cart

import (
	"context"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddItemToCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddItemToCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddItemToCartLogic {
	return AddItemToCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddItemToCartLogic) AddItemToCart(req types.AddItemToCartRequest) (resp *types.AddItemToCartResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
