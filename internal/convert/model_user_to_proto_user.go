package convert

import (
	"github.com/k8scommerce/k8scommerce/services/rpc/user/pb/user"

	"github.com/k8scommerce/k8scommerce/internal/models"
)

func ModelUserToProtoUser(fromModel *models.User, toProto *user.User) {
	// convert the User
	toProto.Id = fromModel.ID
	toProto.FirstName = fromModel.FirstName
	toProto.LastName = fromModel.LastName
	toProto.Email = fromModel.Email
	// toProto.Password = fromModel.Password
}
