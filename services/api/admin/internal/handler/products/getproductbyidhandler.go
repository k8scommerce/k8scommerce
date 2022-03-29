package products

import (
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/logic/products"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProductByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductByIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := products.NewGetProductByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetProductById(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
