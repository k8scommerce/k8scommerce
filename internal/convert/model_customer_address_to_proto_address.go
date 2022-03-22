package convert

import (
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/customer/pb/customer"
)

func ModelCustomerAddressToProtoAddress(fromModel *models.CustomerAddress, toProto *customer.Address) {
	toProto.Street = fromModel.Street
	if fromModel.AptSuite.Valid {
		toProto.AptSuite = fromModel.AptSuite.String
	}
	toProto.City = fromModel.City
	toProto.StateProvince = fromModel.StateProvince
	toProto.PostalCode = fromModel.PostalCode
	toProto.Country = fromModel.Country
	toProto.IsDefault = fromModel.IsDefault
}
