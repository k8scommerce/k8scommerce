package cart

import (
	"context"

	"k8scommerce/internal/utils"
	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBySessionIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBySessionIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBySessionIdLogic {
	return &GetBySessionIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBySessionIdLogic) GetBySessionId(req *types.GetBySessionIdRequest) (resp *types.CartResponse, err error) {
	response, err := l.svcCtx.CartRpc.GetBySessionId(l.ctx, &cart.GetBySessionIdRequest{
		CartId:    req.CartId,
		SessionId: req.SessionId,
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	utils.TransformObj(response, &resp)
	return resp, err
}
