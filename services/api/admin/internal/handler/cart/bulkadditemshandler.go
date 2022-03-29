package cart

import (
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/logic/cart"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BulkAddItemsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BulkAddItemsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := cart.NewBulkAddItemsLogic(r.Context(), svcCtx)
		resp, err := l.BulkAddItems(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
