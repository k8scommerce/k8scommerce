package logic

import (
	"context"
	"net/http"
	"sync"

	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/customer/internal/svc"
	"k8scommerce/services/rpc/customer/pb/customer"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
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
	u := models.Customer{}
	utils.TransformObj(in.Customer, &u)
	if err := l.svcCtx.Repo.Customer().Create(&u); err != nil {
		// logx.Infof("error: %s", err)
		return &customer.CreateCustomerResponse{
			Customer:      nil,
			StatusCode:    http.StatusExpectationFailed,
			StatusMessage: err.Error(),
		}, nil
	}

	// the output object
	out := &customer.Customer{}
	utils.TransformObj(u, &out)

	// the response struct
	return &customer.CreateCustomerResponse{
		Customer:      out,
		StatusCode:    http.StatusOK,
		StatusMessage: "",
	}, nil
}
