package convert

import (
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/customer/pb/customer"
)

func ModelCustomerToProtoCustomer(fromModel *models.Customer, toProto *customer.Customer) {
	// convert the Customer
	toProto.Id = fromModel.ID
	toProto.StoreId = fromModel.StoreID
	toProto.FirstName = fromModel.FirstName
	toProto.LastName = fromModel.LastName
	toProto.Email = fromModel.Email
	toProto.IsVerified = fromModel.IsVerified
}
