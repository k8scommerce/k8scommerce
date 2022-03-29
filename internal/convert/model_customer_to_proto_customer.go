package convert

import (
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/pb/customer"

	"github.com/k8scommerce/k8scommerce/internal/models"
)

func ModelCustomerToProtoCustomer(fromModel *models.Customer, toProto *customer.Customer) {
	// convert the Customer
	toProto.Id = fromModel.ID
	toProto.StoreId = fromModel.StoreID
	toProto.FirstName = fromModel.FirstName
	toProto.LastName = fromModel.LastName
	toProto.Company = fromModel.Company.String
	toProto.Phone = fromModel.Phone.String
	toProto.Email = fromModel.Email
	toProto.IsVerified = fromModel.IsVerified
}
