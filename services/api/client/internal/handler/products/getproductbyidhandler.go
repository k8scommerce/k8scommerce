package products

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"k8scommerce/services/api/client/internal/logic/products"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
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
