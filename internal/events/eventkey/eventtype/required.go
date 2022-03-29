package eventtype

import (
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/pb/customer"
	"github.com/k8scommerce/k8scommerce/services/rpc/store/pb/store"

	"github.com/k8scommerce/k8scommerce/internal/convert"
	"github.com/k8scommerce/k8scommerce/internal/repos"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewRequired() *Required {
	return &Required{
		Customer:     &customer.Customer{},
		Store:        &store.Store{},
		StoreSetting: &store.StoreSetting{},
	}
}

type Required struct {
	Customer     *customer.Customer  `json:"customer"`
	Store        *store.Store        `json:"store"`
	StoreSetting *store.StoreSetting `json:"store_setting"`
}

func (r *Required) Prepare(repo repos.Repo, preloaded *struct {
	StoreId       int64
	CustomerEmail string
	Customer      *customer.Customer
	Store         *store.Store
	StoreSetting  *store.StoreSetting
}) *Required {
	err := mr.Finish(func() error {
		if preloaded.Customer == nil {
			found, err := repo.Customer().GetCustomerByEmail(preloaded.StoreId, preloaded.CustomerEmail)
			if err != nil {
				return status.Errorf(codes.Internal, "could not find customer by id: %s", preloaded.StoreId)
			}

			if found != nil {
				convert.ModelCustomerToProtoCustomer(found, r.Customer)
			}
		} else {
			r.Customer = preloaded.Customer
		}
		return nil
	}, func() error {
		if preloaded.Store == nil {
			found, err := repo.Store().GetStoreById(preloaded.StoreId)
			if err != nil {
				return status.Errorf(codes.Internal, "could not find store by id: %s", preloaded.StoreId)
			}
			if found != nil {
				convert.ModelStoreToProtoStore(found, r.Store)
			}
		} else {
			r.Store = preloaded.Store
		}
		return nil
	}, func() error {
		if preloaded.StoreSetting == nil {
			found, err := repo.StoreSetting().GetStoreSettingById(preloaded.StoreId)
			if err != nil {
				return status.Errorf(codes.Internal, "could not find store setting by store id: %s", preloaded.StoreId)
			}
			if found != nil {
				convert.ModelStoreSettingToProtoStoreSetting(found, r.StoreSetting)
			}
		} else {
			r.StoreSetting = preloaded.StoreSetting
		}
		return nil
	})
	if err != nil {
		logx.Error(status.Errorf(codes.Internal, "could not find required email items in parallel: %s", err.Error()))
	}

	// logx.Infof("REQUIRED: %#v", r)

	return r
}
