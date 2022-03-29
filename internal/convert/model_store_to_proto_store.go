package convert

import (
	"github.com/k8scommerce/k8scommerce/services/rpc/store/pb/store"

	"github.com/k8scommerce/k8scommerce/internal/models"
)

func ModelStoreToProtoStore(fromModel *models.Store, toProto *store.Store) {
	// convert the Store
	toProto.Id = fromModel.ID
	toProto.Name = fromModel.Name
	toProto.Description = fromModel.Description.String
	toProto.Url = fromModel.URL
}
