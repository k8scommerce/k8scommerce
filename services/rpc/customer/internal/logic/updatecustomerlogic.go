package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/customer/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/pb/customer"

	"github.com/k8scommerce/k8scommerce/internal/models"
	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomerLogic {
	return &UpdateCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCustomerLogic) UpdateCustomer(in *customer.UpdateCustomerRequest) (*customer.UpdateCustomerResponse, error) {
	c := models.Customer{}
	utils.TransformObj(in.Customer, &c)
	if err := l.svcCtx.Repo.Customer().Update(&c); err != nil {
		// logx.Infof("error: %s", err)
		return &customer.UpdateCustomerResponse{
			Customer: nil,
		}, nil
	}

	// the output object
	out := &customer.Customer{}
	utils.TransformObj(c, &out)

	// the response struct
	return &customer.UpdateCustomerResponse{
		Customer: out,
	}, nil
}
