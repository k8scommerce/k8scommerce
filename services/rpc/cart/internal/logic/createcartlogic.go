package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/cart/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"

	"github.com/k8scommerce/k8scommerce/internal/convert"
	"github.com/k8scommerce/k8scommerce/internal/groupctx"
	"github.com/k8scommerce/k8scommerce/internal/models"
	"github.com/k8scommerce/k8scommerce/internal/session"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCartLogic {
	return &CreateCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCartLogic) CreateCart(in *cart.CreateCartRequest) (*cart.CartResponse, error) {
	foundCart := models.Cart{
		StoreID: in.StoreId,
	}
	if err := l.svcCtx.Repo.Cart().Create(&foundCart); err != nil {
		logx.Infof("error: %s", err)
		return nil, err
	}

	foundCartItems := []*models.CartItem{}

	ct := &cart.Cart{}
	convert.ModelCartToProtoCart(&foundCart, foundCartItems, ct)

	// add the cart to cache
	l.ctx = groupctx.SetCart(l.ctx, ct)
	res := &cart.CartResponse{}
	err := cache(l.svcCtx.Cache).Get(l.ctx, session.NewSessionId(), groupcache.ProtoSink(res))
	return res, err
}
