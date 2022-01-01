package logic

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/k8scommerce/k8scommerce/pkg/models"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"
	"github.com/k8scommerce/k8scommerce/services/rpc/inventory/inventoryclient"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/mr"
)

type galaxyUpdateItemQuantityInCartLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryUpdateItemQuantityInCartLogic *galaxyUpdateItemQuantityInCartLogicHelper

type UpdateItemQuantityInCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewUpdateItemQuantityInCartLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *UpdateItemQuantityInCartLogic {
	return &UpdateItemQuantityInCartLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *UpdateItemQuantityInCartLogic) UpdateItemQuantityInCart(in *cart.UpdateItemQuantityInCartRequest) (*cart.UpdateItemQuantityInCartResponse, error) {
	var availableQuantity int32 = 0
	res := &cart.UpdateItemQuantityInCartResponse{}

	err := mr.Finish(func() error {
		// check that the cart exists
		result, err := l.svcCtx.Repo.Cart().GetCartByUserId(in.UserId)
		if err != nil {
			return err
		}
		if result == nil {
			return fmt.Errorf("error: no cart exists for user with id: %d", in.UserId)
		}
		return nil
	}, func() error {
		// clear the existing cart
		if entryCartLogic != nil {
			l.mu.Lock()
			err := entryCartLogic.galaxy.Remove(l.ctx, strconv.FormatInt(in.UserId, 10))
			l.mu.Unlock()
			return fmt.Errorf("error: deleting cart cache: %s", err.Error())
		}

		return nil
	}, func() error {
		// check inventory
		result, err := l.svcCtx.InventoryRpc.GetItemQuantity(l.ctx, &inventoryclient.GetItemQuantityRequest{
			Sku:    in.Item.Sku,
			Region: "US",
		})
		if err != nil {
			return err
		}

		// make sure our available quantity is greater than 1
		availableQuantity = result.Quanity
		if availableQuantity == 0 {
			return fmt.Errorf("error: not enough inventory to add to cart")
		}

		return nil
	})

	if err != nil {
		log.Printf("add to cart error: %v", err)
		return nil, err
	}

	// add the item to the cart
	_, err = l.svcCtx.Repo.CartItem().AddItem(
		in.UserId,
		&models.CartItem{
			Sku:       in.Item.Sku,
			Quantity:  int(in.Item.Quantity),
			Price:     in.Item.Price,
			ExpiresAt: in.Item.ExpiresAt.AsTime(),
		})
	if err != nil {
		return nil, err
	}

	cartResponse, cartItems, totalPrice, err := getUpdatedCart(l.svcCtx, in.UserId, res)
	if err != nil {
		return nil, err
	}

	res.Cart = &cart.Cart{
		UserId:     cartResponse.Cart.UserID,
		TotalPrice: totalPrice,
		Items:      cartItems,
	}

	return res, err
}
