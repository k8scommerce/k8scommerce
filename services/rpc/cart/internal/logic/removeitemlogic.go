package logic

import (
	"context"

	"k8scommerce/services/rpc/cart/internal/svc"
	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveItemLogic {
	return &RemoveItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveItemLogic) RemoveItem(in *cart.RemoveItemRequest) (*cart.CartResponse, error) {
	cartId, err := uuid.Parse(in.CartId)
	if err != nil {
		logx.Infof("error: %s", err)
		return nil, err
	}

	err = l.svcCtx.Repo.CartItem().Delete(
		cartId,
		in.Sku,
		false,
	)
	if err != nil {
		return nil, err
	}

	return getNewSessionByCartId(l.ctx, l.svcCtx, cartId)
}
