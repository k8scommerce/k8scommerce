package cart

import (
	"context"

	"k8scommerce/internal/utils"
	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveItemLogic {
	return &RemoveItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveItemLogic) RemoveItem(req *types.RemoveItemRequest) (resp *types.CartResponse, err error) {
	response, err := l.svcCtx.CartRpc.RemoveItem(l.ctx, &cart.RemoveItemRequest{
		CartId: req.CartId,
		Sku:    req.Sku,
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	utils.TransformObj(response, &resp)
	return resp, err
}
