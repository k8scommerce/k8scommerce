package customers

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/customerclient"

	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) SetPasswordLogic {
	return SetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetPasswordLogic) SetPassword(req types.SetPasswordRequest) (resp *types.SetPasswordResponse, err error) {
	resp = &types.SetPasswordResponse{
		Success: false,
	}

	found, err := l.svcCtx.CustomerRpc.SetPassword(l.ctx, &customerclient.SetPasswordRequest{
		StoreId:  l.ctx.Value(types.StoreKey).(int64),
		Code:     req.Code,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	if found.Customer == nil || found.Customer.Id == 0 {
		return resp, nil
	}

	customer := &types.Customer{}
	utils.TransformObj(found.Customer, customer)

	// create the token
	// jwtToken, err := getJwt(
	// 	l.svcCtx.Config.Auth.AccessExpire,
	// 	l.svcCtx.Config.Auth.AccessSecret,
	// 	map[string]interface{}{
	// 		"customerId": found.Customer.Id,
	// 	},
	// )
	// if err != nil {
	// 	return nil, err
	// }

	resp.Customer = types.Customer{
		FirstName:         customer.FirstName,
		LastName:          customer.LastName,
		Email:             customer.Email,
		BillingAddresses:  customer.BillingAddresses,
		ShippingAddresses: customer.ShippingAddresses,
	}

	// resp.JwtToken = *jwtToken
	resp.Success = true

	return resp, err
}
