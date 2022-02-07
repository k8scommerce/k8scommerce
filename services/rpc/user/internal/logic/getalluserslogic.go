package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"k8scommerce/internal/galaxyctx"
	"k8scommerce/services/rpc/user/internal/svc"
	"k8scommerce/services/rpc/user/internal/types"
	"k8scommerce/services/rpc/user/pb/user"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/logx"
)

type galaxyGetAllUsersLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetAllUsersLogic *galaxyGetAllUsersLogicHelper

type GetAllUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetAllUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetAllUsersLogic {
	return &GetAllUsersLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetAllUsersLogic) GetAllUsers(in *user.GetAllUsersRequest) (*user.GetAllUsersResponse, error) {

	// caching goes logic here
	if entryGetAllUsersLogic == nil {
		l.mu.Lock()
		entryGetAllUsersLogic = &galaxyGetAllUsersLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetAllUsersLogic.once.Do(func() {
		fmt.Println(`l.entryGetAllUsersLogic.Do`)

		// register the galaxy one time
		entryGetAllUsersLogic.galaxy = gcache.RegisterGalaxyFunc("GetAllUsers", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {

				found, err := l.svcCtx.Repo.User().GetAllUsers(
					galaxyctx.GetCurrentPage(ctx),
					galaxyctx.GetPageSize(ctx),
					galaxyctx.GetSortOn(ctx),
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
						types.ConvertModelUserToProtoUser(&f, &user)
						users = append(users, &user)
					}
				}

				// the response struct
				item := &user.GetAllUsersResponse{
					Users:        users,
					TotalRecords: totalRecords,
					TotalPages:   totalPages,
				}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &user.GetAllUsersResponse{}

	codec := &galaxycache.ByteCodec{}

	l.ctx = galaxyctx.SetCurrentPage(l.ctx, in.CurrentPage)
	l.ctx = galaxyctx.SetPageSize(l.ctx, in.PageSize)
	l.ctx = galaxyctx.SetSortOn(l.ctx, in.SortOn)

	key := fmt.Sprintf("%d|%d|%s", in.CurrentPage, in.PageSize, in.SortOn)
	if err := entryGetAllUsersLogic.galaxy.Get(l.ctx, key, codec); err != nil {
		return res, nil
	}

	b, err := codec.MarshalBinary()
	if err != nil {
		return res, nil
	}

	err = json.Unmarshal(b, res)
	return res, err

}
