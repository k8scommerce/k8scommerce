package convert

import (
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/catalog/pb/catalog"
)

func ModelProductToProtoProduct(
	modelProduct *models.Product,
	modelVariants *[]models.Variant,
	modelPrices *[]models.Price,
	protoProduct *catalog.Product,
) {
	// convert the product
	protoProduct.Id = modelProduct.ID
	protoProduct.Slug = modelProduct.Slug
	protoProduct.Name = modelProduct.Name
	protoProduct.ShortDescription = modelProduct.ShortDescription.String
	protoProduct.Description = modelProduct.Description.String
	protoProduct.MetaTitle = modelProduct.MetaTitle.String
	protoProduct.MetaDescription = modelProduct.MetaDescription.String
	protoProduct.MetaKeywords = modelProduct.MetaKeywords.String
	protoProduct.Promotionable = modelProduct.Promotionable
	protoProduct.AvailableOn = modelProduct.AvailableOn.Time.String()
	protoProduct.DiscontinueOn = modelProduct.DiscontinueOn.Time.GoString()

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

		protoProduct.Variants = variants
		// protoProduct.Variants[0].
	}
}
