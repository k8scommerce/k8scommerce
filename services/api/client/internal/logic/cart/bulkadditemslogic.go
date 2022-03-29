package cart

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"

	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type BulkAddItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBulkAddItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BulkAddItemsLogic {
	return &BulkAddItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BulkAddItemsLogic) BulkAddItems(req *types.BulkAddItemsRequest) (resp *types.CartResponse, err error) {

	items := []*cart.Item{}
	for _, item := range req.Items {
		items = append(items, &cart.Item{
			Sku:      item.Sku,
			Quantity: item.Quantity,
			Price:    int64(item.Price),
		})
	}

	response, err := l.svcCtx.CartRpc.BulkAddItems(l.ctx, &cart.BulkAddItemsRequest{
		CartId: req.CartId,
		Items:  items,
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	utils.TransformObj(response, &resp)
	return resp, err
}
