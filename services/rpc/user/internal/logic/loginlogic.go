package logic

import (
	"context"

	"k8scommerce/services/rpc/user/internal/svc"
	"k8scommerce/services/rpc/user/pb/user"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *LoginLogic {
	return &LoginLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	found, err := l.svcCtx.Repo.User().Login(in.Email, in.Password)
	if err != nil {
		return &user.LoginResponse{
			User: nil,
		}, nil
	}

	res := &user.LoginResponse{
		User: &user.User{
			Id:        found.ID,
			FirstName: found.FirstName,
			LastName:  found.LastName,
			Email:     found.Email,
		},
	}
	return res, nil
}
