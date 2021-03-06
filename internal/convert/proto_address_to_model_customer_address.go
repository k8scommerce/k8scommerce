package convert

import (
	"database/sql"

	"github.com/k8scommerce/k8scommerce/services/rpc/customer/pb/customer"

	"github.com/k8scommerce/k8scommerce/internal/models"
)

func ProtoAddressToProtoModelCustomerAddress(id *int64, storeId, customerId int64, kind models.AddressKind, fromProto *customer.Address, toModel *models.CustomerAddress) {
	if id != nil {
		toModel.ID = *id
	}
	toModel.StoreID = storeId
	toModel.CustomerID = customerId
	toModel.Kind = kind
	toModel.Street = fromProto.Street
	toModel.AptSuite = sql.NullString{String: fromProto.AptSuite, Valid: true}
	toModel.City = fromProto.City
	toModel.StateProvince = fromProto.StateProvince
	toModel.PostalCode = fromProto.PostalCode
	toModel.Country = fromProto.Country
	toModel.IsDefault = fromProto.IsDefault
}
