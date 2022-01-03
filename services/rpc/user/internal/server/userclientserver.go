// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"k8scommerce/services/rpc/user/internal/logic"
	"k8scommerce/services/rpc/user/internal/svc"
	"k8scommerce/services/rpc/user/pb/user"

	"github.com/localrivet/galaxycache"
)

type UserClientServer struct {
	svcCtx   *svc.ServiceContext
	universe *galaxycache.Universe
}

func NewUserClientServer(svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *UserClientServer {
	return &UserClientServer{
		svcCtx:   svcCtx,
		universe: universe,
	}
}

func (s *UserClientServer) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx, s.universe)
	return l.CreateUser(in)
}

func (s *UserClientServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx, s.universe)
	return l.Login(in)
}
