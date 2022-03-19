package convert

import (
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/catalog/pb/catalog"
)

func ModelCategoryToProtoCategory(modelCategory *models.Category, protoCategory *catalog.Category) {
	// convert the Category
	protoCategory.Id = modelCategory.ID
	protoCategory.ParentId = modelCategory.ParentID.Int64
	protoCategory.Slug = modelCategory.Slug
	protoCategory.Name = modelCategory.Name
	protoCategory.Description = modelCategory.Description.String
	protoCategory.MetaTitle = modelCategory.MetaTitle.String
	protoCategory.MetaDescription = modelCategory.MetaDescription.String
	protoCategory.MetaKeywords = modelCategory.MetaKeywords.String
	protoCategory.SortOrder = modelCategory.SortOrder.Int64
	protoCategory.Depth = modelCategory.Depth.Int64
}