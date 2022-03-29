package users

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/types"
	"github.com/k8scommerce/k8scommerce/services/rpc/user/userclient"

	"github.com/k8scommerce/k8scommerce/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllUsersLogic {
	return GetAllUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUsersLogic) GetAllUsers(req types.GetAllUsersRequest) (resp *types.GetAllUsersResponse, err error) {
	response, err := l.svcCtx.UserRpc.GetAllUsers(l.ctx, &userclient.GetAllUsersRequest{
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
		SortOn:      req.SortOn,
	})
	if err != nil {
		return nil, err
	}

	// convert from one type to another
	// the structs are identical
	out := &types.GetAllUsersResponse{}
	utils.TransformObj(response, &out)
	// b, err := json.Marshal(response)
	// if err != nil {
	// 	return nil, err
	// }
	// err = json.Unmarshal(b, &res)

	return out, err
}
