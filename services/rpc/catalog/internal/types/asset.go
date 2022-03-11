package types

import (
	"k8scommerce/internal/models"
	"k8scommerce/internal/storage/types"
	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/zeromicro/go-zero/core/logx"
)

func ConvertModelAssetToProtoAsset(modelAssets []*models.Asset) (protoAssets []*catalog.Asset) {

	for _, m := range modelAssets {
		var err error
		var sizes *types.ImageSizes

		imageSizes := &types.ImageSizes{}
		sizes, err = imageSizes.Unmarshal(m.Sizes)
		if err != nil {
			logx.Infof("m.Sizes: %s", string(m.Sizes))
			logx.Errorf("error unmarshalling image sizes", err.Error())
		}

		asset := &catalog.Asset{
			Id:          m.ID,
			StoreId:     m.StoreID,
			ProductId:   m.ProductID,
			VariantId:   m.VariantID,
			Name:        m.Name,
			Url:         m.URL,
			ContentType: m.ContentType,
			Kind:        catalog.AssetKind(m.Kind),
		}

		if m.DisplayName.Valid {
			asset.DisplayName = m.DisplayName.String
		}

		if m.SortOrder.Valid {
			asset.SortOrder = m.SortOrder.Int64
		}

		if sizes != nil {
			asset.Sizes = *sizes
		}

		protoAssets = append(protoAssets, asset)
	}

	return protoAssets
}
