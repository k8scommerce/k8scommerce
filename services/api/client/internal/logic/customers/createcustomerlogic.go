package customers

import (
	"context"
	"encoding/json"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
	"k8scommerce/services/rpc/customer/customerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateCustomerLogic {
	return CreateCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCustomerLogic) CreateCustomer(req types.CreateCustomerRequest) (resp *types.CreateCustomerResponse, err error) {
	resp = &types.CreateCustomerResponse{}

	customerObj := &customerclient.CreateCustomerRequest{
		StoreId: l.ctx.Value(types.StoreKey).(int64),
		Customer: &customerclient.Customer{
			StoreId:    l.ctx.Value(types.StoreKey).(int64),
			FirstName:  req.Customer.FirstName,
			LastName:   req.Customer.LastName,
			Email:      req.Customer.Email,
			IsVerified: false,
			BillingAddress: &customerclient.Address{
				Street:        req.Customer.BillingAddress.Street,
				AptSuite:      req.Customer.BillingAddress.AptSuite,
				City:          req.Customer.BillingAddress.City,
				StateProvince: req.Customer.BillingAddress.StateProvince,
				PostalCode:    req.Customer.BillingAddress.PostalCode,
				Country:       req.Customer.BillingAddress.Country,
				IsDefault:     req.Customer.BillingAddress.IsDefault,
			},
		},
	}

	if (types.Address{}) != req.Customer.ShippingAddress {
		customerObj.Customer.ShippingAddress = append(customerObj.Customer.ShippingAddress, &customerclient.Address{
			Street:        req.Customer.ShippingAddress.Street,
			AptSuite:      req.Customer.ShippingAddress.AptSuite,
			City:          req.Customer.ShippingAddress.City,
			StateProvince: req.Customer.ShippingAddress.StateProvince,
			PostalCode:    req.Customer.ShippingAddress.PostalCode,
			Country:       req.Customer.ShippingAddress.Country,
			IsDefault:     req.Customer.ShippingAddress.IsDefault,
		})
	}

	createCustomerResponse, err := l.svcCtx.CustomerRpc.CreateCustomer(l.ctx, customerObj)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(createCustomerResponse)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &resp)

	return resp, err
}
