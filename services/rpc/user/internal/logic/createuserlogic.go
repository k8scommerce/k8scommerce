package logic

import (
	"context"
	"net/http"

	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/user/internal/svc"
	"k8scommerce/services/rpc/user/pb/user"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	u := models.User{}
	utils.TransformObj(in.User, &u)
	if err := l.svcCtx.Repo.User().Create(&u); err != nil {
		// logx.Infof("error: %s", err)
		return &user.CreateUserResponse{
			User:          nil,
			StatusCode:    http.StatusExpectationFailed,
			StatusMessage: err.Error(),
		}, nil
	}

	// the output object
	out := &user.User{}
	utils.TransformObj(u, &out)

	// the response struct
	return &user.CreateUserResponse{
		User:          out,
		StatusCode:    http.StatusOK,
		StatusMessage: "",
	}, nil
}
