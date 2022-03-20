package eventtype

import (
	"k8scommerce/services/rpc/customer/pb/customer"
	"k8scommerce/services/rpc/store/pb/store"
)

type Required struct {
	*customer.Customer
	*store.Store
	*store.StoreSetting
}

func (r *Required) Load() {

}
