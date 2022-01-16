package logic

import (
	"context"

	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateCartItemQuantityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCartItemQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateCartItemQuantityLogic {
	return UpdateCartItemQuantityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCartItemQuantityLogic) UpdateCartItemQuantity(req types.UpdateCartItemQuantityRequest) (resp *types.UpdateCartItemQuantityResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
