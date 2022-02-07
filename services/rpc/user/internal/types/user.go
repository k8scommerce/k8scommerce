package types

import (
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/user/pb/user"
)

func ConvertModelUserToProtoUser(fromModel *models.User, toProto *user.User) {
	// convert the User
	toProto.Id = fromModel.ID
	toProto.FirstName = fromModel.FirstName
	toProto.LastName = fromModel.LastName
	toProto.Email = fromModel.Email
	// toProto.Password = fromModel.Password
}
