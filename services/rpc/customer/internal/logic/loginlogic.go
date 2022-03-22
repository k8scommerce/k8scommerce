package logic

import (
	"context"

	"k8scommerce/services/rpc/customer/internal/svc"
	"k8scommerce/services/rpc/customer/pb/customer"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *LoginLogic {
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

	res := &customer.LoginResponse{
		Customer: &customer.Customer{
			Id:        found.ID,
			FirstName: found.FirstName,
			LastName:  found.LastName,
			Email:     found.Email,
		},
	}
	return res, nil

}
