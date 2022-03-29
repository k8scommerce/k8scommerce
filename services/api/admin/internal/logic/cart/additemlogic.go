package cart

import (
	"context"

	"k8scommerce/internal/utils"
	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddItemLogic {
	return &AddItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddItemLogic) AddItem(req *types.AddItemRequest) (resp *types.CartResponse, err error) {
	response, err := l.svcCtx.CartRpc.AddItem(l.ctx, &cart.AddItemRequest{
		CartId: req.CartId,
		Item: &cart.Item{
			Sku:      req.Item.Sku,
			Quantity: req.Item.Quantity,
			Price:    int64(req.Item.Price),
		},
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	utils.TransformObj(response, &resp)
	return resp, err
}
