package logic

import (
	"context"

	"client/internal/svc"
	"client/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateCartLogic {
	return CreateCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCartLogic) CreateCart(req types.CreateCartRequest) (*types.CreateCartResponse, error) {
	// todo: add your logic here and delete this line

	return &types.CreateCartResponse{}, nil
}
