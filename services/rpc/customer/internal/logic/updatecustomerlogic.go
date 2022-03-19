package logic

import (
	"context"
	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/customer/internal/svc"
	"k8scommerce/services/rpc/customer/pb/customer"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
}

func NewUpdateCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *UpdateCustomerLogic {
	return &UpdateCustomerLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
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