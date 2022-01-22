package logic

import (
	"context"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"
	"strconv"
	"sync"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	mu       sync.Mutex
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *catalog.UpdateCategoryRequest) (*catalog.UpdateCategoryResponse, error) {
	found, err := l.svcCtx.Repo.Category().GetCategoryById(in.Id)
	if err != nil {
		return &catalog.UpdateCategoryResponse{
			// StatusCode:    http.StatusExpectationFailed,
			// StatusMessage: err.Error(),
		}, err
	}

	if err := l.svcCtx.Repo.Category().Update(found); err != nil {
		logx.Infof("error: %s", err)
		return &catalog.UpdateCategoryResponse{
			Category: nil,
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

		// TODO change the removal strategy
		// the issue is we have dynamic key names
		// we can only remove "A" key if we know its name
		// and we don't know all key names
		if entryGetAllCategoriesLogic != nil {
			l.mu.Lock()
			entryGetAllCategoriesLogic.galaxy.Remove(l.ctx, AllCatgoriesKey)
			l.mu.Unlock()
		}

		// same here
		if entryGetProductsByCategoryIdLogic != nil {
			l.mu.Lock()
			entryGetProductsByCategoryIdLogic.galaxy.Remove(l.ctx, AllCatgoriesKey)
			l.mu.Unlock()
		}
	}

	// the output object
	out := &catalog.Category{}
	utils.TransformObj(found, &out)

	// the response struct
	return &catalog.UpdateCategoryResponse{
		Category: out,
		// StatusCode:    http.StatusOK,
		// StatusMessage: "",
	}, err

}
