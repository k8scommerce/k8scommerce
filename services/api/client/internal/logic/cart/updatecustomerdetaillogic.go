package cart

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/pb/customer"

	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCustomerDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCustomerDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomerDetailLogic {
	return &UpdateCustomerDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCustomerDetailLogic) UpdateCustomerDetail(req *types.UpdateCustomerDetailRequest) (resp *types.CartResponse, err error) {

	request := &cart.UpdateCustomerDetailRequest{
		CartId: req.CartId,
	}

	if req.FirstName != "" {
		request.FirstName = req.FirstName
	}
	if req.LastName != "" {
		request.LastName = req.LastName
	}
	if req.Company != "" {
		request.Company = req.Company
	}
	if req.Email != "" {
		request.Email = req.Email
	}
	if req.Phone != "" {
		request.Phone = req.Phone
	}
	if req.BillingAddress != (types.Address{}) {
		request.BillingAddress = &customer.Address{
			Street:        req.BillingAddress.Street,
			AptSuite:      req.BillingAddress.AptSuite,
			City:          req.BillingAddress.City,
			StateProvince: req.BillingAddress.StateProvince,
			PostalCode:    req.BillingAddress.PostalCode,
			Country:       req.BillingAddress.Country,
		}
	}
	if req.ShippingAddress != (types.Address{}) {
		request.ShippingAddress = &customer.Address{
			Street:        req.ShippingAddress.Street,
			AptSuite:      req.ShippingAddress.AptSuite,
			City:          req.ShippingAddress.City,
			StateProvince: req.ShippingAddress.StateProvince,
			PostalCode:    req.ShippingAddress.PostalCode,
			Country:       req.ShippingAddress.Country,
		}
	}

	response, err := l.svcCtx.CartRpc.UpdateCustomerDetail(l.ctx, request)
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	utils.TransformObj(response, &resp)
	return resp, err
}
