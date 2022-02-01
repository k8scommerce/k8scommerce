package Customers

import (
	"context"

	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateCustomerLogic {
	return CreateCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCustomerLogic) CreateCustomer(req types.CreateCustomerRequest) (resp *types.CreateCustomerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
