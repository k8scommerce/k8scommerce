package products

import (
	"net/http"

	"k8scommerce/services/api/client/internal/logic/products"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProductsByCategoryIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProductsByCategoryIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := products.NewGetProductsByCategoryIdLogic(r.Context(), svcCtx)
		resp, err := l.GetProductsByCategoryId(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
