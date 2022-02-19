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
	"github.com/zeromicro/go-zero/core/logx"
)

type galaxyGetAllPermissionGroupsLogicHelper struct {
	once   *sync.Once
	galaxy *galaxycache.Galaxy
}

var entryGetAllPermissionGroupsLogic *galaxyGetAllPermissionGroupsLogicHelper

type GetAllPermissionGroupsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewGetAllPermissionGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *GetAllPermissionGroupsLogic {
	return &GetAllPermissionGroupsLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *GetAllPermissionGroupsLogic) GetAllPermissionGroups(in *user.GetAllPermissionGroupsRequest) (*user.GetAllPermissionGroupsResponse, error) {

	// caching goes logic here
	if entryGetAllPermissionGroupsLogic == nil {
		l.mu.Lock()
		entryGetAllPermissionGroupsLogic = &galaxyGetAllPermissionGroupsLogicHelper{
			once: &sync.Once{},
		}
		l.mu.Unlock()
	}

	entryGetAllPermissionGroupsLogic.once.Do(func() {
		fmt.Println(`l.entryGetAllPermissionGroupsLogic.Do`)

		// register the galaxy one time
		entryGetAllPermissionGroupsLogic.galaxy = gcache.RegisterGalaxyFunc("GetAllPermissionGroups", l.universe, galaxycache.GetterFunc(
			func(ctx context.Context, key string, dest galaxycache.Codec) error {

				found, err := l.svcCtx.Repo.User().GetAllPermissionGroups(
					galaxyctx.GetCurrentPage(ctx),
					galaxyctx.GetPageSize(ctx),
					galaxyctx.GetSortOn(ctx),
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
						types.ConvertModelPermissionGroupToProtoPermissionGroup(&f, &group)
						groups = append(groups, &group)
					}
				}

				// the response struct
				item := &user.GetAllPermissionGroupsResponse{
					PermissionGroups: groups,
					TotalRecords:     totalRecords,
					TotalPages:       totalPages,
				}

				out, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return dest.UnmarshalBinary(out)
			}))
	})

	res := &user.GetAllPermissionGroupsResponse{}

	codec := &galaxycache.ByteCodec{}

	l.ctx = galaxyctx.SetCurrentPage(l.ctx, in.CurrentPage)
	l.ctx = galaxyctx.SetPageSize(l.ctx, in.PageSize)
	l.ctx = galaxyctx.SetSortOn(l.ctx, in.SortOn)

	key := fmt.Sprintf("%d|%d|%s", in.CurrentPage, in.PageSize, in.SortOn)
	if err := entryGetAllPermissionGroupsLogic.galaxy.Get(l.ctx, key, codec); err != nil {
		return res, nil
	}

	b, err := codec.MarshalBinary()
	if err != nil {
		return res, nil
	}

	err = json.Unmarshal(b, res)
	return res, err
}
