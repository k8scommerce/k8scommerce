package cart

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k8scommerce/services/api/admin/internal/logic/cart"
	"k8scommerce/services/api/admin/internal/svc"
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
