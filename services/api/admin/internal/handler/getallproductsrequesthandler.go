package handler

import (
	"net/http"

	"ecomm/services/api/admin/internal/logic"
	"ecomm/services/api/admin/internal/svc"
	"ecomm/services/api/admin/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func getAllProductsRequestHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllProductsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetAllProductsRequestLogic(r.Context(), ctx)
		resp, err := l.GetAllProductsRequest(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
