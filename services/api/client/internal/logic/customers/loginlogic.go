package customers

import (
	"context"
	"k8scommerce/internal/session"
	"k8scommerce/internal/utils"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
	"k8scommerce/services/rpc/customer/customerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.CustomerLoginRequest) (resp *types.CustomerLoginResponse, err error) {
	resp = &types.CustomerLoginResponse{
		Success: false,
	}

	found, err := l.svcCtx.CustomerRpc.Login(l.ctx, &customerclient.LoginRequest{
		StoreId:  l.ctx.Value(types.StoreKey).(int64),
		Email:    req.Email,
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

	// generate the session && jwt tokens
	// sess := l.ctx.Value(types.Session).(session.Session)
	sess := session.NewSession(l.svcCtx.Encrypter, "")
	jwtToken, err := genJwtToken(l.svcCtx, map[string]interface{}{
		"sessionID": sess.GenSessionId(customer.Id),
	})
	if err != nil {
		return nil, err
	}

	resp.Customer = types.Customer{
		FirstName:         customer.FirstName,
		LastName:          customer.LastName,
		Email:             customer.Email,
		BillingAddresses:  customer.BillingAddresses,
		ShippingAddresses: customer.ShippingAddresses,
	}

	resp.JwtToken = *jwtToken
	resp.Success = true

	return resp, nil
}
