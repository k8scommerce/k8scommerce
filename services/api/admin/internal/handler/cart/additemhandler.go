package cart

import (
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/logic/cart"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
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
