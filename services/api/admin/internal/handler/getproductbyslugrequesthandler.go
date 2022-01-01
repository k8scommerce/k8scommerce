package handler

import (
	"net/http"

	"ecomm/services/api/admin/internal/logic"
	"ecomm/services/api/admin/internal/svc"
	"ecomm/services/api/admin/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func getProductBySlugRequestHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductBySlugRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetProductBySlugRequestLogic(r.Context(), ctx)
		resp, err := l.GetProductBySlugRequest(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
