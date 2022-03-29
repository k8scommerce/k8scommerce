package convert

import (
	"encoding/json"

	"github.com/k8scommerce/k8scommerce/services/rpc/store/pb/store"

	"github.com/k8scommerce/k8scommerce/internal/models"

	"github.com/zeromicro/go-zero/core/logx"
)

func ModelStoreSettingToProtoStoreSetting(fromModel *models.StoreSetting, toProto *store.StoreSetting) {
	// convert the Store
	toProto.Id = fromModel.ID
	toProto.StoreId = fromModel.StoreID

	storeConfig := &store.StoreConfig{}
	err := json.Unmarshal(fromModel.Config, storeConfig)
	if err != nil {
		logx.Errorf("unmarshal of store setting config failed: %s", err.Error())
	}

	toProto.Config = storeConfig
}
