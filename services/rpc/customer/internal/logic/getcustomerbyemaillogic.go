package logic

import (
	"context"
	"k8scommerce/internal/convert"
	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/customer/internal/svc"
	"k8scommerce/services/rpc/customer/pb/customer"
	"strings"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	// fetch the addresses in parallel
	err = mr.Finish(func() error {
		addresses := l.getAddressesByKind(out.Id, models.AddressKindBilling)
		out.BillingAddresses = addresses
		return nil
	}, func() error {
		addresses := l.getAddressesByKind(out.Id, models.AddressKindShipping)
		out.ShippingAddresses = addresses
		return nil
	})
	if err != nil {
		logx.Error(status.Errorf(codes.Internal, "could not fetch addresses in parallel: %s", err.Error()))
	}

	// the response struct
	return &customer.GetCustomerByEmailResponse{
		Customer: out,
	}, nil

}

func (l *GetCustomerByEmailLogic) getAddressesByKind(customerId int64, kind models.AddressKind) []*customer.Address {
	addresses, err := l.svcCtx.Repo.CustomerAddress().GetCustomerAddressesByCustomerIdKind(customerId, kind)
	if err != nil {
		logx.Error(status.Errorf(codes.Internal, "could not fetch address kind - %s: %s", kind, err.Error()))
	}
	var addrs []*customer.Address
	for _, addr := range addresses {
		a := &customer.Address{}
		convert.ModelCustomerAddressToProtoAddress(addr, a)
		addrs = append(addrs, a)
	}
	return addrs

}
