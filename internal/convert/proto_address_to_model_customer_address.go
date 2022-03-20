package convert

import (
	"database/sql"
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/customer/pb/customer"
)

func ProtoAddressToProtoModelCustomerAddress(id *int64, storeId, customerId int64, kind models.AddressKind, proto *customer.Address, model *models.CustomerAddress) {
	if id != nil {
		model.ID = *id
	}
	model.StoreID = storeId
	model.CustomerID = customerId
	model.Kind = kind
	model.Street = proto.Street
	model.AptSuite = sql.NullString{String: proto.AptSuite, Valid: true}
	model.City = proto.City
	model.StateProvince = proto.StateProvince
	model.PostalCode = proto.PostalCode
	model.Country = proto.Country
	model.IsDefault = proto.IsDefault
}
