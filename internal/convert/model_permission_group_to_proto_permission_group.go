package convert

import (
	"github.com/k8scommerce/k8scommerce/services/rpc/user/pb/user"

	"github.com/k8scommerce/k8scommerce/internal/models"
)

func ModelPermissionGroupToProtoPermissionGroup(fromModel *models.PermissionGroup, toProto *user.PermissionGroup) {
	// convert the PermissionGroup
	toProto.Id = fromModel.ID
	toProto.GroupName = fromModel.GroupName
}
