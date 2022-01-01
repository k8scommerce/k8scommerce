package handler

import (
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/logic"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func getProductsByCategoryIdRequestHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductsByCategoryIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetProductsByCategoryIdRequestLogic(r.Context(), ctx)
		resp, err := l.GetProductsByCategoryIdRequest(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}