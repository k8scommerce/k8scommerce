package products

import (
	"net/http"

	"k8scommerce/services/api/admin/internal/logic/products"
	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateProductRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := products.NewCreateProductLogic(r.Context(), svcCtx)
		resp, err := l.CreateProduct(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
