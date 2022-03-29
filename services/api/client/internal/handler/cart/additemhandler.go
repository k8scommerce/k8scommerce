package cart

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k8scommerce/services/api/client/internal/logic/cart"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
)

func AddItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddItemRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := cart.NewAddItemLogic(r.Context(), svcCtx)
		resp, err := l.AddItem(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
