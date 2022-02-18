package customers

import (
	"context"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustomerLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCustomerLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) CustomerLoginLogic {
	return CustomerLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CustomerLoginLogic) CustomerLogin(req types.CustomerLoginRequest) (resp *types.CustomerLoginResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
