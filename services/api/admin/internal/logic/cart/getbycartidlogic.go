package cart

import (
	"context"

	"k8scommerce/internal/utils"
	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByCartIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByCartIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByCartIdLogic {
	return &GetByCartIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByCartIdLogic) GetByCartId(req *types.GetByCartIdRequest) (resp *types.CartResponse, err error) {
	response, err := l.svcCtx.CartRpc.GetByCartId(l.ctx, &cart.GetByCartIdRequest{
		CartId: req.CartId,
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	utils.TransformObj(response, &resp)
	return resp, err
}
