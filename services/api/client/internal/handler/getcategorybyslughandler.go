package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"k8scommerce/services/api/client/internal/logic"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
)

func getCategoryBySlugHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCategoryBySlugRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetCategoryBySlugLogic(r.Context(), ctx)
		resp, err := l.GetCategoryBySlug(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}