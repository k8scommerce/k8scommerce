package convert

import (
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/customer/pb/customer"
)

func ModelCustomerAddressToProtoAddress(model *models.CustomerAddress, proto *customer.Address) {
	proto.Street = model.Street
	if model.AptSuite.Valid {
		proto.AptSuite = model.AptSuite.String
	}
	proto.City = model.City
	proto.StateProvince = model.StateProvince
	proto.PostalCode = model.PostalCode
	proto.Country = model.Country
	proto.IsDefault = model.IsDefault
}
