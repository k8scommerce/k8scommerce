package handler

import (
	"net/http"

	"client/internal/logic"
	"client/internal/svc"
	"client/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func createCartHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateCartRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateCartLogic(r.Context(), ctx)
		resp, err := l.CreateCart(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
