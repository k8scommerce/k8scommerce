package types

import (
	"k8scommerce/internal/models"
	"k8scommerce/internal/storage/types"
	"k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/zeromicro/go-zero/core/logx"
)

func ConvertModelAssetToProtoAsset(modelAssets *[]models.Asset) (protoAssets []*catalog.Asset) {
	if modelAssets != nil {
		for _, v := range *modelAssets {
			imageSizes := &types.ImageSizes{}
			is, err := imageSizes.Unmarshal(v.Sizes)
			if err != nil {
				logx.Errorf("error unmarshalling image sizes", err.Error())
			}

			asset := &catalog.Asset{
				Id:          v.ID,
				StoreId:     v.StoreID,
				ProductId:   v.ProductID,
				VariantId:   v.VariantID,
				Name:        v.Name,
				DisplayName: v.DisplayName.String,
				Url:         v.URL,
				ContentType: v.ContentType,
				SortOrder:   v.SortOrder.Int64,
				Kind:        stringToProtoAssetKind(v.Kind),
			}

			for _, size := range is.Sizes {
				asset.Sizes = append(asset.Sizes, &catalog.ImageSize{
					Tag: size.Tag,
					URL: size.URL,
				})
			}

			protoAssets = append(protoAssets, asset)
		}
	}
	return protoAssets
}

func stringToProtoAssetKind(assetKind string) catalog.AssetKind {
	switch assetKind {
	case catalog.AssetKind_image.String():
		return catalog.AssetKind_image
	case catalog.AssetKind_document.String():
		return catalog.AssetKind_document
	case catalog.AssetKind_audio.String():
		return catalog.AssetKind_audio
	case catalog.AssetKind_video.String():
		return catalog.AssetKind_video
	case catalog.AssetKind_archive.String():
		return catalog.AssetKind_archive
	default:
		return catalog.AssetKind_unknown
	}
}
