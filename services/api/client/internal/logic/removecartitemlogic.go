package logic

import (
	"context"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveCartItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveCartItemLogic {
	return RemoveCartItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveCartItemLogic) RemoveCartItem(req types.RemoveCartItemRequest) (resp *types.RemoveCartItemResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
