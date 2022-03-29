package convert

import (
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/k8scommerce/k8scommerce/internal/models"
)

func ModelProductToProtoProduct(
	modelProduct *models.Product,
	modelVariants *[]models.Variant,
	modelPrices *[]models.Price,
	toProto *catalog.Product,
) {
	// convert the product
	toProto.Id = modelProduct.ID
	toProto.Slug = modelProduct.Slug
	toProto.Name = modelProduct.Name
	toProto.ShortDescription = modelProduct.ShortDescription.String
	toProto.Description = modelProduct.Description.String
	toProto.MetaTitle = modelProduct.MetaTitle.String
	toProto.MetaDescription = modelProduct.MetaDescription.String
	toProto.MetaKeywords = modelProduct.MetaKeywords.String
	toProto.Promotionable = modelProduct.Promotionable
	toProto.AvailableOn = modelProduct.AvailableOn.Time.String()
	toProto.DiscontinueOn = modelProduct.DiscontinueOn.Time.GoString()

	if modelVariants != nil {
		// convert the variants
		variants := []*catalog.Variant{}

		for _, v := range *modelVariants {
			variant := &catalog.Variant{
				Id:                 v.ID,
				ProductId:          v.ProductID,
				IsDefault:          v.IsDefault,
				Sku:                v.Sku,
				SortOrder:          int32(v.SortOrder),
				CostAmount:         v.CostAmount.Int64,
				CostCurrency:       v.CostCurrency.String,
				TrackInventory:     v.TrackInventory,
				TaxCategoryId:      v.TaxCategoryID.Int64,
				ShippingCategoryId: v.ShippingCategoryID.Int64,
				DiscontinueOn:      v.DiscontinueOn.Time.String(),
				Weight:             v.Weight.Float64,
				Height:             v.Height.Float64,
				Width:              v.Width.Float64,
				Depth:              v.Depth.Float64,
			}

			if modelPrices != nil {
				for _, modelPrice := range *modelPrices {
					if modelPrice.VariantID == variant.Id {
						variant.Price = &catalog.Price{
							Id:          modelPrice.ID,
							VariantId:   modelPrice.VariantID,
							SalePrice:   modelPrice.SalePrice,
							RetailPrice: modelPrice.RetailPrice.Int64,
							Currency:    modelPrice.Currency.String,
						}
					}
				}
			}

			variants = append(variants, variant)

		}

		toProto.Variants = variants
		// toProto.Variants[0].
	}
}
