package convert

import (
	"encoding/json"
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/cart/pb/cart"
	"k8scommerce/services/rpc/customer/pb/customer"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelCartToProtoCart(fromCartModel *models.Cart, fromCartModelItems []*models.CartItem, toProto *cart.Cart) {
	// convert the Cart
	toProto.Id = fromCartModel.ID.String()
	toProto.StoreId = fromCartModel.StoreID
	toProto.FirstName = fromCartModel.FirstName.String
	toProto.LastName = fromCartModel.LastName.String
	toProto.Company = fromCartModel.Company.String
	toProto.Phone = fromCartModel.Phone.String
	toProto.Email = fromCartModel.Email.String
	toProto.Status = cart.CartStatus(fromCartModel.Status)

	if fromCartModel.Currency.Valid {
		toProto.Currency = fromCartModel.Currency.String
	}

	if fromCartModel.DiscountID.Valid {
		toProto.DiscountId = fromCartModel.DiscountID.Int64
	}

	if fromCartModel.AbandonedAt.Valid {
		toProto.AbandonedAt = timestamppb.New(fromCartModel.AbandonedAt.Time)
	}

	if fromCartModel.ExpiresAt.Valid {
		toProto.ExpiresAt = timestamppb.New(fromCartModel.ExpiresAt.Time)
	}

	{
		billingAddress := &customer.Address{}
		err := json.Unmarshal(fromCartModel.BillingAddress, billingAddress)
		if err != nil {
			logx.Errorf("unmarshal of billing address failed: %s", err.Error())
		}
		toProto.BillingAddress = billingAddress
	}

	{
		shippingAddress := &customer.Address{}
		err := json.Unmarshal(fromCartModel.ShippingAddress, shippingAddress)
		if err != nil {
			logx.Errorf("unmarshal of shipping address failed: %s", err.Error())
		}
		toProto.ShippingAddress = shippingAddress
	}

	var totalPrice int64 = 0
	for _, item := range fromCartModelItems {
		out := &cart.Item{}
		ModelCartItemToProtoCartItem(item, out)
		toProto.Items = append(toProto.Items, out)
		totalPrice += item.Price
	}
	toProto.TotalPrice = totalPrice
}
