package logic

import (
	"context"
	"net/http"
	"sync"

	"github.com/k8scommerce/k8scommerce/pkg/models"
	"github.com/k8scommerce/k8scommerce/pkg/utils"
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/pb/customer"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
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
