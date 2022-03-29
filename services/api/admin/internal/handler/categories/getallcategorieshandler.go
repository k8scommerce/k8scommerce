package categories

import (
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/logic/categories"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAllCategoriesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := categories.NewGetAllCategoriesLogic(r.Context(), svcCtx)
		resp, err := l.GetAllCategories()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
