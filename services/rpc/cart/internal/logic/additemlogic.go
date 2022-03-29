package logic

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"k8scommerce/internal/models"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/cart/internal/svc"
	"k8scommerce/services/rpc/cart/pb/cart"
	"k8scommerce/services/rpc/inventory/inventoryclient"
	"k8scommerce/services/rpc/othersbought/othersboughtclient"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddItemLogic {
	return &AddItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddItemLogic) AddItem(in *cart.AddItemRequest) (*cart.CartResponse, error) {

	response := &cart.CartResponse{
		SimilarProducts: &cart.SimilarProducts{},
		OthersBought:    &cart.OthersBought{},
	}

	cartId, err := uuid.Parse(in.CartId)
	if err != nil {
		logx.Infof("error: %s", err)
		return nil, err
	}

	foundCart := &models.Cart{}

	err = mr.Finish(func() error {
		// check that the cart exists
		// the cart doesn't exist, we need to reload from the database
		foundCart, err = l.svcCtx.Repo.Cart().GetByCartId(cartId)
		if err != nil {
			logx.Infof("error: %s", err)
			return status.Errorf(codes.NotFound, "could not find a cart by id of: %s", cartId)
		}
		if foundCart == nil {
			return status.Errorf(codes.NotFound, "error: no cart exists for user with id: %s", cartId)
		}

		return nil
	}, func() error {
		// check inventory
		result, err := l.svcCtx.InventoryRpc.GetItemQuantity(l.ctx, &inventoryclient.GetItemQuantityRequest{
			Sku:     in.Item.Sku,
			StoreId: foundCart.StoreID,
		})
		if err != nil {
			return err
		}

		// make sure our available quantity is greater than 1
		if result.StockLevel.Quantity < 1 {
			return fmt.Errorf("error: not enough inventory to add to cart")
		}

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
		response.OthersBought.Variants = sp.Variants
		return
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
		// response.SimilarProducts.Variants = sp.Variants

		return nil
	})

	if err != nil {
		log.Printf("add to cart error: %v", err)
		return nil, err
	}

	quantity := in.Item.Quantity
	if quantity == 0 {
		quantity = 1
	}

	// add the item to the cart
	item := &models.CartItem{
		Sku:      in.Item.Sku,
		Quantity: int(quantity),
		Price:    in.Item.Price,
	}

	if in.Item.ExpiresAt.IsValid() {
		item.ExpiresAt = sql.NullTime{
			Time:  in.Item.ExpiresAt.AsTime(),
			Valid: true,
		}
	}

	_, err = l.svcCtx.Repo.CartItem().AddItem(cartId, item)
	if err != nil {
		return nil, err
	}

	return getNewSessionByCartId(l.ctx, l.svcCtx, cartId)
}
