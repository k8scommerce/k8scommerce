package convert

import (
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/cart/pb/cart"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelCartItemToProtoCartItem(fromCartItemModel *models.CartItem, toProto *cart.Item) {
	if fromCartItemModel.ExpiresAt.Valid {
		toProto.ExpiresAt = timestamppb.New(fromCartItemModel.ExpiresAt.Time)
	}

	toProto.Sku = fromCartItemModel.Sku
	toProto.Price = fromCartItemModel.Price
	toProto.Quantity = int32(fromCartItemModel.Quantity)
}
