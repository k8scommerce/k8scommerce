package cart

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"

	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type AttachCustomerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttachCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttachCustomerLogic {
	return &AttachCustomerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttachCustomerLogic) AttachCustomer(req *types.AttachCustomerRequest) (resp *types.CartResponse, err error) {
	response, err := l.svcCtx.CartRpc.AttachCustomer(l.ctx, &cart.AttachCustomerRequest{
		CartId:        req.CartId,
		CustomerEmail: req.CustomerEmail,
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	utils.TransformObj(response, &resp)
	return resp, err
}
