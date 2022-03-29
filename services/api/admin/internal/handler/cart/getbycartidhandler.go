package cart

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k8scommerce/services/api/admin/internal/logic/cart"
	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
)

func GetByCartIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetByCartIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := cart.NewGetByCartIdLogic(r.Context(), svcCtx)
		resp, err := l.GetByCartId(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
