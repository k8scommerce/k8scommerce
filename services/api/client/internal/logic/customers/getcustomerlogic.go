package customers

import (
	"context"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCustomerLogic {
	return GetCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCustomerLogic) GetCustomer() (resp *types.GetCustomerResponse, err error) {
	resp = &types.GetCustomerResponse{}

	// response, err := l.svcCtx.CustomerRpc.GetCustomerByEmail(l.ctx, &customerclient.SetPasswordRequest{
	// 	StoreId:  l.ctx.Value(types.StoreKey).(int64),
	// 	Code:     req.Code,
	// 	Password: req.Password,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// utils.TransformObj(response, &resp)

	return resp, err
}
