package cart

import (
	"context"

	"k8scommerce/internal/utils"
	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCartLogic {
	return &CreateCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCartLogic) CreateCart() (resp *types.CartResponse, err error) {
	response, err := l.svcCtx.CartRpc.CreateCart(l.ctx, &cart.CreateCartRequest{
		StoreId: l.ctx.Value(types.StoreKey).(int64),
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	utils.TransformObj(response, &resp)
	return resp, err
}
