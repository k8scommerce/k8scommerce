package logic

import (
	"context"
	"database/sql"
	"encoding/json"

	"k8scommerce/services/rpc/cart/internal/svc"
	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCustomerDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCustomerDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomerDetailLogic {
	return &UpdateCustomerDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCustomerDetailLogic) UpdateCustomerDetail(in *cart.UpdateCustomerDetailRequest) (*cart.CartResponse, error) {
	cartId, _ := uuid.Parse(in.CartId)
	var err error
	foundCart, err := l.svcCtx.Repo.Cart().GetByCartId(cartId)
	if err != nil {
		logx.Infof("error: %s", err)
		// we either have a valid cart or we don't
		return nil, err
	}

	if in.FirstName != "" {
		foundCart.FirstName = sql.NullString{String: in.FirstName, Valid: true}
	}
	if in.LastName != "" {
		foundCart.LastName = sql.NullString{String: in.LastName, Valid: true}
	}
	if in.Email != "" {
		foundCart.Email = sql.NullString{String: in.Email, Valid: true}
	}
	if in.Phone != "" {
		foundCart.Phone = sql.NullString{String: in.Phone, Valid: true}
	}
	if in.Company != "" {
		foundCart.Company = sql.NullString{String: in.Company, Valid: true}
	}
	if in.BillingAddress != nil {
		out, _ := json.Marshal(in.BillingAddress)
		foundCart.BillingAddress = out
	}
	if in.ShippingAddress != nil {
		out, _ := json.Marshal(in.ShippingAddress)
		foundCart.ShippingAddress = out
	}

	err = l.svcCtx.Repo.Cart().Update(foundCart)
	if err != nil {
		logx.Infof("error: %s", err)
	}

	return getNewSessionByCartId(l.ctx, l.svcCtx, cartId)
}
