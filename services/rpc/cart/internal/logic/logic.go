package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/k8scommerce/k8scommerce/services/rpc/cart/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"

	"github.com/k8scommerce/k8scommerce/internal/convert"
	"github.com/k8scommerce/k8scommerce/internal/gcache"
	"github.com/k8scommerce/k8scommerce/internal/groupctx"
	"github.com/k8scommerce/k8scommerce/internal/models"
	"github.com/k8scommerce/k8scommerce/internal/session"

	"github.com/google/uuid"
	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

const Group_Cart = "Cart"

func cache(cache gcache.Cache) *groupcache.Group {
	return cache.NewGroup(Group_Cart, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, sessionId string, dest groupcache.Sink) error {
			ct := groupctx.GetCart(ctx)
			if ct == nil {
				return fmt.Errorf("could not find cart in context")
			}

			// Set the groupcache to expire after 24 hours
			if err := dest.SetProto(&cart.CartResponse{
				Cart:      ct,
				SessionId: sessionId,
			}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}

func getNewSessionByCartId(ctx context.Context, svcCtx *svc.ServiceContext, cartId uuid.UUID) (*cart.CartResponse, error) {
	var err error
	foundCart := &models.Cart{}
	foundCartItems := []*models.CartItem{}
	err = mr.Finish(func() error {
		foundCart, err = svcCtx.Repo.Cart().GetByCartId(cartId)
		if err != nil {
			logx.Infof("error: %s", err)
			return err
		}
		return nil
	}, func() error {
		foundCartItems, err = svcCtx.Repo.CartItem().GetByCartId(cartId)
		if err != nil {
			logx.Infof("error: %s", err)
		}
		return nil
	})
	if err != nil {
		// we either have a valid cart or we don't
		return nil, err
	}

	ct := &cart.Cart{}
	convert.ModelCartToProtoCart(foundCart, foundCartItems, ct)

	// add the cart to cache
	ctx = groupctx.SetCart(ctx, ct)
	res := &cart.CartResponse{}
	err = cache(svcCtx.Cache).Get(ctx, session.NewSessionId(), groupcache.ProtoSink(res))
	return res, err
}
