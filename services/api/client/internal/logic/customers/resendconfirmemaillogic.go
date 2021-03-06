package customers

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/customerclient"

	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResendConfirmEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResendConfirmEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) ResendConfirmEmailLogic {
	return ResendConfirmEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResendConfirmEmailLogic) ResendConfirmEmail(req types.ResendConfirmEmailRequest) (resp *types.ResendConfirmEmailResponse, err error) {
	resp = &types.ResendConfirmEmailResponse{
		Success: false,
	}

	response, err := l.svcCtx.CustomerRpc.SendConfirmEmailAddressEmail(l.ctx, &customerclient.SendConfirmEmailAddressEmailRequest{
		StoreId: l.ctx.Value(types.StoreKey).(int64),
		Email:   req.Email,
	})
	if err != nil {
		return nil, err
	}

	utils.TransformObj(response, &resp)

	return resp, err
}
