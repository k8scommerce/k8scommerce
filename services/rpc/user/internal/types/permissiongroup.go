package types

import (
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/user/pb/user"
)

func ConvertModelPermissionGroupToProtoPermissionGroup(fromModel *models.PermissionGroup, toProto *user.PermissionGroup) {
	// convert the PermissionGroup
	toProto.Id = fromModel.ID
	toProto.GroupName = fromModel.GroupName
}
