package assets

import (
	"net/http"

	"k8scommerce/services/api/admin/internal/logic/assets"
	"k8scommerce/services/api/admin/internal/svc"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := assets.NewUploadLogic(r, r.Context(), svcCtx)
		resp, err := l.Upload()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
