package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"k8scommerce/services/api/client/internal/logic"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
)

func clearCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClearCartRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewClearCartLogic(r.Context(), svcCtx)
		resp, err := l.ClearCart(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}