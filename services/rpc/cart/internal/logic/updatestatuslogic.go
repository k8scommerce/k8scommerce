package logic

import (
	"context"
	"database/sql"
	"time"

	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/cart/internal/svc"
	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStatusLogic {
	return &UpdateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStatusLogic) UpdateStatus(in *cart.UpdateStatusRequest) (*cart.CartResponse, error) {
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

	foundCart.Status = models.CartStatus(in.Status)
	if foundCart.Status == models.CartStatusAbandoned {
		foundCart.AbandonedAt = sql.NullTime{Time: time.Now(), Valid: true}
	}

	err = l.svcCtx.Repo.Cart().Update(foundCart)
	if err != nil {
		logx.Infof("error: %s", err)
		return nil, err
	}

	return getNewSessionByCartId(l.ctx, l.svcCtx, cartId)
}
