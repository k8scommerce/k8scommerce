package products

import (
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/logic/products"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProductBySkuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductBySkuRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := products.NewGetProductBySkuLogic(r.Context(), svcCtx)
		resp, err := l.GetProductBySku(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
