package logic

import (
	"context"

	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/customer/internal/svc"
	"k8scommerce/services/rpc/customer/pb/customer"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *customer.LoginRequest) (*customer.LoginResponse, error) {
	found, err := l.svcCtx.Repo.Customer().Login(in.StoreId, in.Email, in.Password)
	if err != nil {
		return &customer.LoginResponse{}, nil
	}

	out := &customer.Customer{}
	utils.TransformObj(found, &out)

	// fetch the addresses in parallel
	err = mr.Finish(func() error {
		addresses := getAddressesByKind(l.svcCtx.Repo, out.Id, models.AddressKindBilling)
		out.BillingAddresses = addresses
		return nil
	}, func() error {
		addresses := getAddressesByKind(l.svcCtx.Repo, out.Id, models.AddressKindShipping)
		out.ShippingAddresses = addresses
		return nil
	})
	if err != nil {
		logx.Error(status.Errorf(codes.Internal, "could not fetch addresses in parallel: %s", err.Error()))
	}

	res := &customer.LoginResponse{
		Customer: out,
	}
	return res, nil

}
