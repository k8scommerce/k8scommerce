package cart

import (
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/logic/cart"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AttachCustomerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AttachCustomerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := cart.NewAttachCustomerLogic(r.Context(), svcCtx)
		resp, err := l.AttachCustomer(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}