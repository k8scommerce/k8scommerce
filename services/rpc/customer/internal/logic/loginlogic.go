package logic

import (
	"context"
	"net/http"
	"sync"

	"k8scommerce/services/rpc/customer/internal/svc"
	"k8scommerce/services/rpc/customer/pb/customer"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyLoginLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryLoginLogic *galaxyLoginLogicHelper

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *LoginLogic {
	return &LoginLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *LoginLogic) Login(in *customer.LoginRequest) (*customer.LoginResponse, error) {

	found, err := l.svcCtx.Repo.Customer().Login(in.Email, in.Password)
	if err != nil {
		return &customer.LoginResponse{
			Customer:      nil,
			StatusCode:    http.StatusExpectationFailed,
			StatusMessage: err.Error(),
		}, nil
	}

	res := &customer.LoginResponse{
		Customer: &customer.Customer{
			Id:        found.ID,
			FirstName: found.FirstName,
			LastName:  found.LastName,
			Email:     found.Email,
		},
		StatusCode:    http.StatusOK,
		StatusMessage: "",
	}
	return res, nil

}
