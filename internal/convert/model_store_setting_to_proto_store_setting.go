package convert

import (
	"encoding/json"
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/store/pb/store"

	"github.com/zeromicro/go-zero/core/logx"
)

func ModelStoreSettingToProtoStoreSetting(fromModel *models.StoreSetting, toProto *store.StoreSetting) {
	// convert the Store
	toProto.Id = fromModel.ID
	toProto.StoreId = fromModel.StoreID

	err := json.Unmarshal(fromModel.Config, toProto.Config)
	if err != nil {
		logx.Errorf("unmarshal of store setting config failed: %s", err.Error())
	}
}