package users

import (
	"context"
	"encoding/json"

	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
	"k8scommerce/services/rpc/user/userclient"

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
	resp = &types.GetAllUsersResponse{}
	b, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &resp)

	return resp, err
}
