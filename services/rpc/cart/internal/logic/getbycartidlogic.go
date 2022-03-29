package logic

import (
	"context"

	"k8scommerce/services/rpc/cart/internal/svc"
	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

const Group_GetByCartId = "GetByCartId"

type GetByCartIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByCartIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByCartIdLogic {
	return &GetByCartIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByCartIdLogic) GetByCartId(in *cart.GetByCartIdRequest) (*cart.CartResponse, error) {
	cartId, err := uuid.Parse(in.CartId)
	if err != nil {
		logx.Infof("error: %s", err)
		return nil, err
	}

	return getNewSessionByCartId(l.ctx, l.svcCtx, cartId)
}
