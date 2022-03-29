package logic

import (
	"context"
	"time"

	"github.com/k8scommerce/k8scommerce/services/rpc/user/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/user/pb/user"

	"github.com/k8scommerce/k8scommerce/internal/convert"
	"github.com/k8scommerce/k8scommerce/internal/gcache"
	"github.com/k8scommerce/k8scommerce/internal/groupctx"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

const Group_GetAllUsers = "GetAllUsers"

var Group_GetAllUsersKey = func(currentPage, pageSize int64, sortOn string) string {
	return gcache.ToKey(Group_GetAllUsers, currentPage, pageSize, sortOn)
}

type GetAllUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUsersLogic {
	return &GetAllUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllUsersLogic) GetAllUsers(in *user.GetAllUsersRequest) (*user.GetAllUsersResponse, error) {
	l.ctx = groupctx.SetCurrentPage(l.ctx, in.CurrentPage)
	l.ctx = groupctx.SetPageSize(l.ctx, in.PageSize)
	l.ctx = groupctx.SetSortOn(l.ctx, in.SortOn)
	res := &user.GetAllUsersResponse{}
	err := l.cache().Get(l.ctx, Group_GetAllUsersKey(in.CurrentPage, in.PageSize, in.SortOn), groupcache.ProtoSink(res))
	return res, err

}

func (l *GetAllUsersLogic) cache() *groupcache.Group {
	return l.svcCtx.Cache.NewGroup(Group_GetAllUsers, 128<<20, groupcache.GetterFunc(
		func(ctx context.Context, id string, dest groupcache.Sink) error {
			found, err := l.svcCtx.Repo.User().GetAllUsers(
				groupctx.GetCurrentPage(ctx),
				groupctx.GetPageSize(ctx),
				groupctx.GetSortOn(ctx),
			)
			if err != nil {
				logx.Infof("error: %s", err)
				return err
			}

			users := []*user.User{}

			var totalRecords int64 = 0
			var totalPages int64 = 0

			if found != nil {
				totalRecords = found.PagingStats.TotalRecords
				totalPages = found.PagingStats.TotalPages

				for _, f := range found.Users {
					user := user.User{}
					convert.ModelUserToProtoUser(&f, &user)
					users = append(users, &user)
				}
			}

			// Set the groupcache to expire after 24 hours
			if err := dest.SetProto(&user.GetAllUsersResponse{
				Users:        users,
				TotalRecords: totalRecords,
				TotalPages:   totalPages,
			}, time.Now().Add(time.Hour*24)); err != nil {
				return err
			}
			return nil
		},
	))
}
