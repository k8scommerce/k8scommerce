package customers

import (
	"context"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
	"k8scommerce/services/rpc/customer/customerclient"

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

	// create the token
	jwtToken, err := getJwt(
		l.svcCtx.Config.Auth.AccessExpire,
		l.svcCtx.Config.Auth.AccessSecret,
		map[string]interface{}{
			"customerId": found.Customer.Id,
		},
	)
	if err != nil {
		return nil, err
	}

	customer := types.Customer{
		FirstName: found.Customer.FirstName,
		LastName:  found.Customer.LastName,
		Email:     found.Customer.Email,
	}

	resp.JwtToken = *jwtToken
	resp.Customer = customer
	resp.Success = true

	return resp, err
}
