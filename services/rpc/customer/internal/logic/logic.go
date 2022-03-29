package logic

import (
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/pb/customer"
	"github.com/k8scommerce/k8scommerce/services/rpc/store/pb/store"

	"github.com/k8scommerce/k8scommerce/internal/convert"
	"github.com/k8scommerce/k8scommerce/internal/models"
	"github.com/k8scommerce/k8scommerce/internal/repos"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getProtoStoreByStoreId(repo repos.Repo, storeId int64) (*store.Store, error) {
	foundStore, err := repo.Store().GetStoreById(storeId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not find store by id: %d", storeId)
	}

	protoStore := &store.Store{}
	if foundStore != nil {
		convert.ModelStoreToProtoStore(foundStore, protoStore)
	}

	return protoStore, nil
}

func getAddressesByKind(repo repos.Repo, customerId int64, kind models.AddressKind) []*customer.Address {
	addresses, err := repo.CustomerAddress().GetCustomerAddressesByCustomerIdKind(customerId, kind)
	if err != nil {
		logx.Error(status.Errorf(codes.Internal, "could not fetch address kind - %s: %s", kind, err.Error()))
	}
	var addrs []*customer.Address
	for _, addr := range addresses {
		a := &customer.Address{}
		convert.ModelCustomerAddressToProtoAddress(addr, a)
		addrs = append(addrs, a)
	}
	return addrs
}
