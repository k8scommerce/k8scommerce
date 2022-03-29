package logic

import (
	"context"
	"time"

	"k8scommerce/internal/convert"
	"k8scommerce/internal/gcache"
	"k8scommerce/internal/groupctx"
	"k8scommerce/services/rpc/user/internal/svc"
	"k8scommerce/services/rpc/user/pb/user"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

const Group_GetAllPermissionGroups = "GetAllPermissionGroups"

var Group_GetAllPermissionGroupsKey = func(currentPage, pageSize int64, sortOn string) string {
	return gcache.ToKey(Group_GetAllPermissionGroups, currentPage, pageSize, sortOn)
}

type GetAllPermissionGroupsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllPermissionGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPermissionGroupsLogic {
	return &GetAllPermissionGroupsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllPermissionGroupsLogic) GetAllPermissionGroups(in *user.GetAllPermissionGroupsRequest) (*user.GetAllPermissionGroupsResponse, error) {
	l.ctx = groupctx.SetCurrentPage(l.ctx, in.CurrentPage)
	l.ctx = groupctx.SetPageSize(l.ctx, in.PageSize)
	l.ctx = groupctx.SetSortOn(l.ctx, in.SortOn)
	res := &user.GetAllPermissionGroupsResponse{}
	err := l.cache().Get(l.ctx, Group_GetAllPermissionGroupsKey(in.CurrentPage, in.PageSize, in.SortOn), groupcache.ProtoSink(res))
	return res, err
}

func (l *GetAllPermissionGroupsLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetAllPermissionGroups, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.User().GetAllPermissionGroups(
				groupctx.GetCurrentPage(ctx),
				groupctx.GetPageSize(ctx),
				groupctx.GetSortOn(ctx),
			)
			if err != nil {
				logx.Infof("error: %s", err)
				return err
			}

			groups := []*user.PermissionGroup{}

			var totalRecords int64 = 0
			var totalPages int64 = 0

			if found != nil {
				totalRecords = found.PagingStats.TotalRecords
				totalPages = found.PagingStats.TotalPages

				for _, f := range found.PermissionGroups {
					group := user.PermissionGroup{}
					convert.ModelPermissionGroupToProtoPermissionGroup(&f, &group)
					groups = append(groups, &group)
				}
			}

			// Set the groupcache to expire after 24 hours
			if err := dest.SetProto(&user.GetAllPermissionGroupsResponse{
				PermissionGroups: groups,
				TotalRecords:     totalRecords,
				TotalPages:       totalPages,
			}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
