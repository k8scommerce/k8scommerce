package logic

import (
	"context"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/customer/internal/svc"
	"k8scommerce/services/rpc/customer/pb/customer"
	"strings"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomerByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCustomerByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetCustomerByEmailLogic {
	return &GetCustomerByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCustomerByEmailLogic) GetCustomerByEmail(in *customer.GetCustomerByEmailRequest) (*customer.GetCustomerByEmailResponse, error) {
	found, err := l.svcCtx.Repo.Customer().GetCustomerByEmail(in.StoreId, in.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "sql: no rows in result set") {
			return &customer.GetCustomerByEmailResponse{
				Customer: nil,
			}, err
		}
	}

	out := &customer.Customer{}
	utils.TransformObj(found, &out)

	if found != nil {
		out.IsVerified = found.IsVerified.Bool
	}

	// the response struct
	return &customer.GetCustomerByEmailResponse{
		Customer: out,
	}, nil

}
