package logic

import (
	"context"
	"database/sql"
	"encoding/json"

	"k8scommerce/internal/convert"
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/cart/internal/svc"
	"k8scommerce/services/rpc/cart/pb/cart"
	"k8scommerce/services/rpc/customer/pb/customer"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type AttachCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAttachCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttachCustomerLogic {
	return &AttachCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AttachCustomerLogic) AttachCustomer(in *cart.AttachCustomerRequest) (*cart.CartResponse, error) {
	cartId, _ := uuid.Parse(in.CartId)
	foundCart, err := l.svcCtx.Repo.Cart().GetByCartId(cartId)
	if err != nil {
		logx.Infof("error: %s", err)
	}

	foundCustomer, err := l.svcCtx.Repo.Customer().GetCustomerByEmail(foundCart.StoreID, in.CustomerEmail)
	if err != nil {
		logx.Infof("error: %s", err)
		return nil, err
	}

	if foundCustomer == nil {
		logx.Infof("no customer found: %s", err)
		return nil, err
	}

	// get addresses and cartitems in parallel
	mr.Finish(func() error {
		addresses, _ := l.svcCtx.Repo.CustomerAddress().GetCustomerAddressesByCustomerIdKind(foundCustomer.ID, models.AddressKindBilling)
		for _, address := range addresses {
			if address.IsDefault {
				toProto := &customer.Address{}
				convert.ModelCustomerAddressToProtoAddress(address, toProto)
				out, _ := json.Marshal(toProto)
				foundCart.BillingAddress = out
			}
		}
		return nil
	}, func() error {
		addresses, _ := l.svcCtx.Repo.CustomerAddress().GetCustomerAddressesByCustomerIdKind(foundCustomer.ID, models.AddressKindShipping)
		for _, address := range addresses {
			if address.IsDefault {
				toProto := &customer.Address{}
				convert.ModelCustomerAddressToProtoAddress(address, toProto)
				out, _ := json.Marshal(toProto)
				foundCart.ShippingAddress = out
			}
		}
		return nil
	})

	if foundCustomer.FirstName != "" {
		foundCart.FirstName = sql.NullString{String: foundCustomer.FirstName, Valid: true}
	}
	if foundCustomer.LastName != "" {
		foundCart.LastName = sql.NullString{String: foundCustomer.LastName, Valid: true}
	}
	if foundCustomer.Email != "" {
		foundCart.Email = sql.NullString{String: foundCustomer.Email, Valid: true}
	}
	foundCart.Phone = foundCustomer.Phone
	foundCart.Company = foundCustomer.Company

	// override any customer data

	err = l.svcCtx.Repo.Cart().Update(foundCart)
	if err != nil {
		logx.Infof("error: %s", err)
	}

	return getNewSessionByCartId(l.ctx, l.svcCtx, cartId)
}
