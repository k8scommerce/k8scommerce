package logic

import (
	"context"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"
	"strconv"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(in *catalog.DeleteCategoryRequest) (*catalog.DeleteCategoryResponse, error) {
	_, err := l.svcCtx.Repo.Category().GetCategoryById(in.Id)
	if err != nil {
		return &catalog.DeleteCategoryResponse{}, err
	}

	// delete the Category
	if err := l.svcCtx.Repo.Category().Delete(in.Id); err != nil {
		return &catalog.DeleteCategoryResponse{
			// StatusCode:    http.StatusExpectationFailed,
			// StatusMessage: err.Error(),
		}, err
	}

	// invalidate the cache for this record
	{
		if entryGetCategoryByIdLogic != nil {
			l.mu.Lock()
			entryGetCategoryByIdLogic.galaxy.Remove(l.ctx, strconv.Itoa(int(in.Id)))
			l.mu.Unlock()
		}
		if entryGetAllCategoriesLogic != nil {
			l.mu.Lock()
			entryGetAllCategoriesLogic.galaxy.Remove(l.ctx, AllCatgoriesKey)
			l.mu.Unlock()
		}
	}

	// the response struct
	return &catalog.DeleteCategoryResponse{
		// StatusCode:    http.StatusOK,
		// StatusMessage: "",
	}, nil
}
