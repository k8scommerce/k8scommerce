package logic

import (
	"context"

	"fmt"
	"log"
	"strconv"
	"sync"

	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/cart/internal/svc"
	"k8scommerce/services/rpc/cart/pb/cart"
	"k8scommerce/services/rpc/inventory/inventoryclient"
	"k8scommerce/services/rpc/othersbought/othersboughtclient"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type AddItemToCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewAddItemToCartLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *AddItemToCartLogic {
	return &AddItemToCartLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *AddItemToCartLogic) AddItemToCart(in *cart.AddItemToCartRequest) (*cart.AddItemToCartResponse, error) {
	var availableQuantity int32 = 0
	res := &cart.AddItemToCartResponse{
		SimilarProducts: &cart.SimilarProducts{},
		OthersBought:    &cart.OthersBought{},
	}

	err := mr.Finish(func() error {
		// check that the cart exists
		result, err := l.svcCtx.Repo.Cart().GetCartByCustomerId(in.CustomerId)
		if err != nil {
			return err
		}
		if result == nil {
			return fmt.Errorf("error: no cart exists for user with id: %d", in.CustomerId)
		}
		return nil
	}, func() error {
		// clear the existing cart
		if entryCartLogic != nil {
			l.mu.Lock()
			err := entryCartLogic.galaxy.Remove(l.ctx, strconv.FormatInt(in.CustomerId, 10))
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
	}, func() error {
		// // similar products
		// result, err := l.svcCtx.SimilarProductsRpc.GetSimilarProductsBySku(l.ctx, &similarproductsclient.GetSimilarProductsBySkuRequest{
		// 	Sku: in.Item.Sku,
		// })
		// if err != nil {
		// 	return err
		// }

		// sp := &cart.SimilarProducts{}
		// err = utils.TransformObj(result.Variants, &sp.Variants)
		// if err != nil {
		// 	return fmt.Errorf("error: SimilarProducts error: %s", err.Error())
		// 	// return err
		// }
		// res.SimilarProducts.Variants = sp.Variants

		return nil
	}, func() (err error) {
		// others bought
		result, err := l.svcCtx.OtherBoughtRpc.GetOthersBoughtBySku(l.ctx, &othersboughtclient.GetOthersBoughtBySkuRequest{
			Sku: in.Item.Sku,
		})
		if err != nil {
			return err
		}

		sp := &cart.OthersBought{}
		err = utils.TransformObj(result.Variants, &sp.Variants)
		if err != nil {
			return fmt.Errorf("error: OthersBought error: %s", err.Error())
		}
		res.OthersBought.Variants = sp.Variants
		return
	})

	if err != nil {
		log.Printf("add to cart error: %v", err)
		return nil, err
	}

	// add the item to the cart
	_, err = l.svcCtx.Repo.CartItem().AddItem(
		in.CustomerId,
		&models.CartItem{
			Sku:       in.Item.Sku,
			Quantity:  int(in.Item.Quantity),
			Price:     in.Item.Price,
			ExpiresAt: in.Item.ExpiresAt.AsTime(),
		})
	if err != nil {
		return nil, err
	}

	cartResponse, cartItems, totalPrice, err := getUpdatedCart(l.svcCtx, in.CustomerId, res)
	if err != nil {
		return nil, err
	}

	res.Cart = &cart.Cart{
		CustomerId: cartResponse.Cart.CustomerID,
		TotalPrice: totalPrice,
		Items:      cartItems,
	}

	return res, err
}

// func (l *AddItemToCartLogic) notify() {

// }
