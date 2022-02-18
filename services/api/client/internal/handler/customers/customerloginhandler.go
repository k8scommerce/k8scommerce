package customers

import (
	"net/http"

	"k8scommerce/services/api/client/internal/logic/customers"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CustomerLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CustomerLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := customers.NewCustomerLoginLogic(r.Context(), svcCtx)
		resp, err := l.CustomerLogin(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
