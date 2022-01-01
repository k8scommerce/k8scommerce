package logic

import (
	"context"
	"net/http"

	"github.com/k8s-commerce/k8s-commerce/services/rpc/user/internal/svc"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/user/user"

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
	found, err := l.svcCtx.Repo.User().Login(in.Username, in.Password)
	if err != nil {
		return &user.LoginResponse{
			User:          nil,
			StatusCode:    http.StatusExpectationFailed,
			StatusMessage: err.Error(),
		}, nil
	}

	res := &user.LoginResponse{
		User: &user.User{
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
