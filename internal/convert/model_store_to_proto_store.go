package convert

import (
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/store/pb/store"
)

func ModelStoreToProtoStore(modelStore *models.Store, protoStore *store.Store) {
	// convert the Store
	protoStore.Id = modelStore.ID
	protoStore.Name = modelStore.Name
	protoStore.Description = modelStore.Description.String
	protoStore.Url = modelStore.URL
}
