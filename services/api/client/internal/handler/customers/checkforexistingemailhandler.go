package customers

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k8scommerce/services/api/client/internal/logic/customers"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
)

func CheckForExistingEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckForExistingEmailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := customers.NewCheckForExistingEmailLogic(r.Context(), svcCtx)
		resp, err := l.CheckForExistingEmail(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
