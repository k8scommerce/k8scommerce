package cart

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"

	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateItemQuantityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateItemQuantityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateItemQuantityLogic {
	return &UpdateItemQuantityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateItemQuantityLogic) UpdateItemQuantity(req *types.UpdateItemQuantityRequest) (resp *types.CartResponse, err error) {
	response, err := l.svcCtx.CartRpc.UpdateItemQuantity(l.ctx, &cart.UpdateItemQuantityRequest{
		CartId:   req.CartId,
		Sku:      req.Sku,
		Quantity: req.Quantity,
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	utils.TransformObj(response, &resp)
	return resp, err
}
