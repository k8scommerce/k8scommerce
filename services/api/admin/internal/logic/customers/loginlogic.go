package customers

import (
	"context"
	"k8scommerce/internal/utils"
	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
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

	res, err := l.svcCtx.CustomerRpc.Login(l.ctx, &customerclient.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	if res.Customer == nil {
		return resp, nil
	}

	// create the token

	customer := types.Customer{}
	utils.TransformObj(res.Customer, &customer)

	resp.Customer = customer
	resp.Success = true

	return resp, nil
}
