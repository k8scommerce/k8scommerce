package logic

import (
	"context"
	"sync"

	"k8scommerce/internal/convert"
	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/customer/internal/svc"
	"k8scommerce/services/rpc/customer/pb/customer"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type galaxyCreateCustomerLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryCreateCustomerLogic *galaxyCreateCustomerLogicHelper

type CreateCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewCreateCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *CreateCustomerLogic {
	return &CreateCustomerLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *CreateCustomerLogic) CreateCustomer(in *customer.CreateCustomerRequest) (*customer.CreateCustomerResponse, error) {
	c := models.Customer{}
	utils.TransformObj(in.Customer, &c)

	if err := l.svcCtx.Repo.Customer().Create(&c); err != nil {
		return &customer.CreateCustomerResponse{
			Customer: nil,
		}, status.Errorf(codes.AlreadyExists, "could not create customer: %s", err.Error())
	}

	// the output object
	out := &customer.Customer{}
	utils.TransformObj(c, &out)

	// save the addressses in parallel
	err := mr.Finish(func() error {
		if in.Customer.BillingAddresses != nil {
			for _, addr := range in.Customer.BillingAddresses {
				address, err := l.saveCustomerAddress(in.StoreId, out.Id, addr, models.AddressKindBilling)
				if err != nil {
					logx.Error(status.Errorf(codes.Internal, "could not create billing address: %s", err.Error()))
				}
				out.BillingAddresses = append(out.BillingAddresses, address)
			}
		}
		return nil
	}, func() error {
		if in.Customer.ShippingAddresses != nil {
			for _, addr := range in.Customer.ShippingAddresses {
				address, err := l.saveCustomerAddress(in.StoreId, out.Id, addr, models.AddressKindShipping)
				if err != nil {
					logx.Error(status.Errorf(codes.Internal, "could not create shipping address: %s", err.Error()))
				}
				out.ShippingAddresses = append(out.ShippingAddresses, address)
			}
		}
		return nil
	})
	if err != nil {
		logx.Error(status.Errorf(codes.Internal, "could not save addresses in parallel: %s", err.Error()))
	}

	// the response struct
	return &customer.CreateCustomerResponse{
		Customer: out,
	}, nil
}

func (l *CreateCustomerLogic) saveCustomerAddress(storeId, customerId int64, protoAddress *customer.Address, kind models.AddressKind) (*customer.Address, error) {
	modelAddress := models.CustomerAddress{}
	convert.ProtoAddressToProtoModelCustomerAddress(nil, storeId, customerId, kind, protoAddress, &modelAddress)
	if err := l.svcCtx.Repo.CustomerAddress().Create(&modelAddress); err != nil {
		return nil, status.Errorf(codes.Internal, "could not create customer address: %s", err.Error())
	}
	utils.TransformObj(modelAddress, protoAddress)
	return protoAddress, nil
}
