package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/cart/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"

	"github.com/google/uuid"
	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetBySessionIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBySessionIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBySessionIdLogic {
	return &GetBySessionIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBySessionIdLogic) GetBySessionId(in *cart.GetBySessionIdRequest) (*cart.CartResponse, error) {
	res := &cart.CartResponse{}
	err := cache(l.svcCtx.Cache).Get(l.ctx, in.SessionId, groupcache.ProtoSink(res))
	if err != nil {
		cartId, err := uuid.Parse(in.CartId)
		if err != nil {
			logx.Infof("error: %s", err)
			return nil, err
		}

		return getNewSessionByCartId(l.ctx, l.svcCtx, cartId)
	}

	return res, nil
}
