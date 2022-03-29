package logic

import (
	"context"

	"k8scommerce/services/rpc/user/internal/svc"
	"k8scommerce/services/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	found, err := l.svcCtx.Repo.User().Login(in.Email, in.Password)
	if err != nil {
		return &user.LoginResponse{}, nil
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
