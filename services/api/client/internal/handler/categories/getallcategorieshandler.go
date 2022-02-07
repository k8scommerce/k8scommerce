package categories

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"k8scommerce/services/api/client/internal/logic/categories"
	"k8scommerce/services/api/client/internal/svc"
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
