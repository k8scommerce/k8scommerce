package customers

import (
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/logic/customers"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ResendConfirmEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ResendConfirmEmailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := customers.NewResendConfirmEmailLogic(r.Context(), svcCtx)
		resp, err := l.ResendConfirmEmail(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
