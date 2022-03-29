package convert

import (
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/k8scommerce/k8scommerce/internal/models"
)

func ModelCategoryToProtoCategory(fromModel *models.Category, toProto *catalog.Category) {
	// convert the Category
	toProto.Id = fromModel.ID
	toProto.ParentId = fromModel.ParentID.Int64
	toProto.Slug = fromModel.Slug
	toProto.Name = fromModel.Name
	toProto.Description = fromModel.Description.String
	toProto.MetaTitle = fromModel.MetaTitle.String
	toProto.MetaDescription = fromModel.MetaDescription.String
	toProto.MetaKeywords = fromModel.MetaKeywords.String
	toProto.SortOrder = fromModel.SortOrder.Int64
	toProto.Depth = fromModel.Depth.Int64
}
