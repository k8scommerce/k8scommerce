package logic

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/k8scommerce/k8scommerce/services/rpc/cart/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"
	"github.com/k8scommerce/k8scommerce/services/rpc/inventory/inventoryclient"
	"github.com/k8scommerce/k8scommerce/services/rpc/inventory/pb/inventory"

	"github.com/k8scommerce/k8scommerce/internal/convert"
	"github.com/k8scommerce/k8scommerce/internal/groupctx"
	"github.com/k8scommerce/k8scommerce/internal/models"
	"github.com/k8scommerce/k8scommerce/internal/session"

	"github.com/google/uuid"
	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BulkAddItemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBulkAddItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BulkAddItemsLogic {
	return &BulkAddItemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BulkAddItemsLogic) BulkAddItems(in *cart.BulkAddItemsRequest) (*cart.CartResponse, error) {
	response := &cart.CartResponse{
		SimilarProducts: &cart.SimilarProducts{},
		OthersBought:    &cart.OthersBought{},
	}

	cartId, err := uuid.Parse(in.CartId)
	if err != nil {
		logx.Infof("error: %s", err)
		return nil, err
	}

	foundCartItems := []*models.CartItem{}
	foundCart := &models.Cart{}
	var stockLevels = map[string]*inventory.StockLevel{}
	var notes = map[string]string{}

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
		var skus []string
		for _, item := range in.Items {
			skus = append(skus, item.Sku)
		}

		result, err := l.svcCtx.InventoryRpc.GetItemsQuantity(l.ctx, &inventoryclient.GetItemsQuantityRequest{
			Skus:    skus,
			StoreId: foundCart.StoreID,
		})
		if err != nil {
			return err
		}

		// array to map
		for _, stockLevel := range result.StockLevels {
			stockLevels[stockLevel.Sku] = stockLevel
		}

		return nil
	})

	if err != nil {
		log.Printf("add to cart error: %v", err)
		return nil, err
	}

	for _, item := range in.Items {
		quantity := item.Quantity
		if quantity == 0 {
			quantity = 1
		}
		notes[item.Sku] = ""
		if stockLevels[item.Sku].Quantity < 1 {
			quantity = stockLevels[item.Sku].Quantity
			notes[item.Sku] = fmt.Sprintf("%s is out of stock", item.Sku)
		} else if item.Quantity >= stockLevels[item.Sku].Quantity {
			quantity = stockLevels[item.Sku].Quantity
			notes[item.Sku] = fmt.Sprintf("%s has only %d remaining", item.Sku, stockLevels[item.Sku].Quantity)
		}

		// add the item to the cart
		cartItem := &models.CartItem{
			Sku:      item.Sku,
			Quantity: int(quantity),
			Price:    item.Price,
		}

		if item.ExpiresAt.IsValid() {
			cartItem.ExpiresAt = sql.NullTime{
				Time:  cartItem.ExpiresAt.Time,
				Valid: true,
			}
		}

		// each sql iteration returns the same items
		updatedCartItem, err := l.svcCtx.Repo.CartItem().AddItem(cartId, cartItem)
		if err != nil {
			return nil, err
		}

		foundCartItems = append(foundCartItems, updatedCartItem)
	}

	ct := &cart.Cart{}
	convert.ModelCartToProtoCart(foundCart, foundCartItems, ct)

	for _, cartItem := range ct.Items {
		cartItem.Note = notes[cartItem.Sku]
	}

	// add the cart to cache
	l.ctx = groupctx.SetCart(l.ctx, ct)
	res := &cart.CartResponse{
		SimilarProducts: response.SimilarProducts,
		OthersBought:    response.OthersBought,
	}
	err = cache(l.svcCtx.Cache).Get(l.ctx, session.NewSessionId(), groupcache.ProtoSink(res))
	return res, err
}
