package Categories

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"k8scommerce/services/api/admin/internal/logic/Categories"
	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
)

func DeleteCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCategoryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := Categories.NewDeleteCategoryLogic(r.Context(), svcCtx)
		resp, err := l.DeleteCategory(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
