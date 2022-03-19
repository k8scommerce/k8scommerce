package logic

import (
	"k8scommerce/internal/convert"
	"k8scommerce/internal/repos"
	"k8scommerce/services/rpc/store/pb/store"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getProtoStoreByStoreId(repo repos.Repo, storeId int64) (*store.Store, error) {
	foundStore, err := repo.Store().GetStoreById(storeId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not find store by id: %s", storeId)
	}

	protoStore := &store.Store{}
	if foundStore != nil {
		convert.ModelStoreToProtoStore(foundStore, protoStore)
	}

	return protoStore, nil
}
