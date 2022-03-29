package customers

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/customerclient"

	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForgotPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewForgotPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) ForgotPasswordLogic {
	return ForgotPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ForgotPasswordLogic) ForgotPassword(req types.ForgotPasswordRequest) (resp *types.ForgotPasswordResponse, err error) {
	resp = &types.ForgotPasswordResponse{
		Success: false,
	}

	response, err := l.svcCtx.CustomerRpc.SendForgotPasswordEmail(l.ctx, &customerclient.SendForgotPasswordEmailRequest{
		StoreId: l.ctx.Value(types.StoreKey).(int64),
		Email:   req.Email,
	})
	if err != nil {
		return nil, err
	}

	utils.TransformObj(response, &resp)

	return resp, err
}
