package cart

import (
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/logic/cart"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cart.NewCreateCartLogic(r.Context(), svcCtx)
		resp, err := l.CreateCart()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
