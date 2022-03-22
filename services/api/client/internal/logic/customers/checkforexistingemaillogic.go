package customers

import (
	"context"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
	"k8scommerce/services/rpc/customer/customerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckForExistingEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckForExistingEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckForExistingEmailLogic {
	return CheckForExistingEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckForExistingEmailLogic) CheckForExistingEmail(req types.CheckForExistingEmailRequest) (resp *types.CheckForExistingEmailResponse, err error) {
	resp = &types.CheckForExistingEmailResponse{
		Exists: false,
	}
	found, err := l.svcCtx.CustomerRpc.GetCustomerByEmail(l.ctx, &customerclient.GetCustomerByEmailRequest{
		StoreId: l.ctx.Value(types.StoreKey).(int64),
		Email:   req.Email,
	})
	if err != nil {
		return nil, err
	}

	if found.Customer != nil {
		resp.Exists = true
		resp.IsVerified = found.Customer.IsVerified
	}
	return resp, err
}
