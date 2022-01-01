package logic

import (
	"context"

	"ecomm/services/api/client/internal/svc"
	"ecomm/services/api/client/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddToCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddToCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddToCartLogic {
	return AddToCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddToCartLogic) AddToCart(req types.AddToCartRequest) (*types.AddToCartResponse, error) {
	// todo: add your logic here and delete this line

	return &types.AddToCartResponse{
		Cart: types.Cart{
			UserId: 1,
		},
	}, nil
}
