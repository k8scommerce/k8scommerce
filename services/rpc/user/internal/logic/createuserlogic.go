package logic

import (
	"context"

	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/user/internal/svc"
	"k8scommerce/services/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	u := models.User{}
	utils.TransformObj(in.User, &u)
	if err := l.svcCtx.Repo.User().Create(&u); err != nil {
		// logx.Infof("error: %s", err)
		return &user.CreateUserResponse{
			User: nil,
		}, nil
	}

	// the output object
	out := &user.User{}
	utils.TransformObj(u, &out)

	// the response struct
	return &user.CreateUserResponse{
		User: out,
	}, nil
}
