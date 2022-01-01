package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/cart/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveItemInCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
}

func NewRemoveItemInCartLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *RemoveItemInCartLogic {
	return &RemoveItemInCartLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *RemoveItemInCartLogic) RemoveItemInCart(in *cart.RemoveItemInCartRequest) (*cart.RemoveItemInCartResponse, error) {
	err := l.svcCtx.Repo.CartItem().Delete(
		in.UserId,
		in.Sku,
		false,
	)
	if err != nil {
		return nil, err
	}

	res := &cart.RemoveItemInCartResponse{}
	cartResponse, cartItems, totalPrice, err := getUpdatedCart(l.svcCtx, in.UserId, res)
	if err != nil {
		return nil, err
	}

	res.Cart = &cart.Cart{
		UserId:     cartResponse.Cart.UserID,
		TotalPrice: totalPrice,
		Items:      cartItems,
	}

	return res, nil
}
