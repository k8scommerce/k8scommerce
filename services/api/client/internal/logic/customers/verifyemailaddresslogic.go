package customers

import (
	"context"
	"encoding/json"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
	"k8scommerce/services/rpc/customer/customerclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyEmailAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyEmailAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyEmailAddressLogic {
	return VerifyEmailAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyEmailAddressLogic) VerifyEmailAddress(req types.VerifyEmailAddressRequest) (resp *types.VerifyEmailAddressResponse, err error) {
	resp = &types.VerifyEmailAddressResponse{
		Success: false,
	}

	response, err := l.svcCtx.CustomerRpc.VerifyEmailAddress(l.ctx, &customerclient.VerifyEmailAddressRequest{
		StoreId: l.ctx.Value(types.StoreKey).(int64),
		Code:    req.Code,
	})
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &resp)

	return resp, err
}
