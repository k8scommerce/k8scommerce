package handler

import (
	"net/http"

	"ecomm/services/api/client/internal/logic"
	"ecomm/services/api/client/internal/svc"
	"ecomm/services/api/client/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func getProductByIdRequestHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductByIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetProductByIdRequestLogic(r.Context(), ctx)
		resp, err := l.GetProductByIdRequest(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
