package logic

import (
	"context"
	"database/sql"
	"sync"

	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/customer/internal/svc"
	"k8scommerce/services/rpc/customer/pb/customer"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
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

	if in.Customer.IsVerified {
		c.IsVerified = sql.NullBool{Bool: true, Valid: true}
	} else {
		c.IsVerified = sql.NullBool{Bool: false, Valid: true}
	}

	logx.Infof("INCOMING CUSTOMER: %#v", in.Customer)
	logx.Infof("TRANSFORMED CUSTOMER: %#v", c)

	if err := l.svcCtx.Repo.Customer().Create(&c); err != nil {
		return &customer.CreateCustomerResponse{
			Customer: nil,
		}, status.Errorf(codes.Internal, "could not create customer: %s", err.Error())
	}

	// if in.Customer.BillingAddress != nil {
	// 	if in.Customer.BillingAddress.IsDefault {
	// 		c.IsVerified = sql.NullBool{Bool: true, Valid: true}
	// 	}
	// }

	// the output object
	out := &customer.Customer{}
	utils.TransformObj(c, &out)

	// the response struct
	return &customer.CreateCustomerResponse{
		Customer: out,
	}, nil
}
