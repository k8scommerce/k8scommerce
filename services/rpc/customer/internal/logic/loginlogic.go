package logic

import (
	"context"
	"sync"

	"github.com/k8s-commerce/k8s-commerce/services/rpc/customer/internal/svc"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/customer/pb/customer"

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

	res := &customer.LoginResponse{}
	return res, nil
}
