package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/cart/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type ClearCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClearCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearCartLogic {
	return &ClearCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClearCartLogic) ClearCart(in *cart.ClearCartRequest) (*cart.CartResponse, error) {
	cartId, err := uuid.Parse(in.CartId)
	if err != nil {
		logx.Infof("error: %s", err)
		return nil, err
	}

	foundCart, err := l.svcCtx.Repo.Cart().GetByCartId(cartId)
	if err != nil {
		logx.Infof("error: %s", err)
		return nil, err
	}

	err = l.svcCtx.Repo.CartItem().ClearItems(foundCart.ID, false)
	if err != nil {
		logx.Infof("error: %s", err)
	}

	return getNewSessionByCartId(l.ctx, l.svcCtx, cartId)
}
