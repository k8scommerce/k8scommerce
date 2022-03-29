package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/cart/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateItemQuantityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateItemQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateItemQuantityLogic {
	return &UpdateItemQuantityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateItemQuantityLogic) UpdateItemQuantity(in *cart.UpdateItemQuantityRequest) (*cart.CartResponse, error) {
	cartId, err := uuid.Parse(in.CartId)
	if err != nil {
		logx.Infof("error: %s", err)
		return nil, err
	}

	_, err = l.svcCtx.Repo.CartItem().UpdateQuantity(cartId, in.Sku, int(in.Quantity))
	if err != nil {
		return nil, err
	}

	return getNewSessionByCartId(l.ctx, l.svcCtx, cartId)

}
