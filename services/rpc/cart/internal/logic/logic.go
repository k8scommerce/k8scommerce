package logic

import (
	"github.com/k8s-commerce/k8s-commerce/services/rpc/cart/internal/repos"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/cart/internal/svc"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/cart/pb/cart"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func getUpdatedCart(svcCtx *svc.ServiceContext, userId int64, res interface{}) (
	cartResponse *repos.CartResponse,
	cartItems []*cart.Item,
	totalPrice int64,
	err error,
) {
	// get the whole cart
	cartResponse, err = svcCtx.Repo.Cart().GetCartByUserId(userId)
	if err != nil {
		return nil, nil, totalPrice, err
	}

	for _, item := range cartResponse.Items {
		cartItems = append(cartItems, &cart.Item{
			UserId:    item.UserID,
			Sku:       item.Sku,
			Quantity:  int32(item.Quantity),
			ExpiresAt: timestamppb.New(item.ExpiresAt),
		})
	}

	return cartResponse, cartItems, totalPrice, nil
}

// func transformVariants(variants []*cart.Variant) {
// 	for _, v := range variants {
// 		sp.Variants = append(sp.Variants, &cart.Variant{
// 			Id:                 v.Id,
// 			ProductId:          v.ProductId,
// 			IsDefault:          v.IsDefault,
// 			Sku:                v.Sku,
// 			SortOrder:          v.SortOrder,
// 			CostAmount:         v.CostAmount,
// 			CostCurrency:       v.CostCurrency,
// 			TrackInventory:     v.TrackInventory,
// 			TaxCategoryId:      v.TaxCategoryId,
// 			ShippingCategoryId: v.ShippingCategoryId,
// 			DiscontinueOn:      v.DiscontinueOn,
// 			Weight:             v.Weight,
// 			Height:             v.Height,
// 			Width:              v.Width,
// 			Depth:              v.Depth,
// 			Price: &cart.Price{
// 				Id:              v.Price.Id,
// 				VariantId:       v.Id,
// 				Amount:          v.Price.Amount,
// 				CompareAtAmount: v.Price.CompareAtAmount,
// 				Currency:        v.Price.Currency,
// 			},
// 		})
// 	}
// }
// for _, v := range result.Variants {
// 	sp.Variants = append(sp.Variants, &cart.Variant{
// 		Id:                 v.Id,
// 		ProductId:          v.ProductId,
// 		IsDefault:          v.IsDefault,
// 		Sku:                v.Sku,
// 		SortOrder:          v.SortOrder,
// 		CostAmount:         v.CostAmount,
// 		CostCurrency:       v.CostCurrency,
// 		TrackInventory:     v.TrackInventory,
// 		TaxCategoryId:      v.TaxCategoryId,
// 		ShippingCategoryId: v.ShippingCategoryId,
// 		DiscontinueOn:      v.DiscontinueOn,
// 		Weight:             v.Weight,
// 		Height:             v.Height,
// 		Width:              v.Width,
// 		Depth:              v.Depth,
// 		Price: &cart.Price{
// 			Id:              v.Price.Id,
// 			VariantId:       v.Id,
// 			Amount:          v.Price.Amount,
// 			CompareAtAmount: v.Price.CompareAtAmount,
// 			Currency:        v.Price.Currency,
// 		},
// 	})
// }
