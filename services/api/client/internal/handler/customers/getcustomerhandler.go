package customers

import (
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/logic/customers"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCustomerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := customers.NewGetCustomerLogic(r.Context(), svcCtx)
		resp, err := l.GetCustomer()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
