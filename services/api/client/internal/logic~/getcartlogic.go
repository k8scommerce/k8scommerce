package logic

import (
	"context"

	"client/internal/svc"
	"client/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCartLogic {
	return GetCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCartLogic) GetCart(req types.GetCartRequest) (*types.GetCartResponse, error) {
	// todo: add your logic here and delete this line

	return &types.GetCartResponse{}, nil
}
