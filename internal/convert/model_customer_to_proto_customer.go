package convert

import (
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/customer/pb/customer"
)

func ModelCustomerToProtoCustomer(modelCustomer *models.Customer, protoCustomer *customer.Customer) {
	// convert the Customer
	protoCustomer.Id = modelCustomer.ID
	protoCustomer.StoreId = modelCustomer.StoreID
	protoCustomer.FirstName = modelCustomer.FirstName
	protoCustomer.LastName = modelCustomer.LastName
	protoCustomer.Email = modelCustomer.Email
	protoCustomer.IsVerified = modelCustomer.IsVerified
}
