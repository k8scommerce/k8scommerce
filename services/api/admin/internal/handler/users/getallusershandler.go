package users

import (
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/logic/users"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAllUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllUsersRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := users.NewGetAllUsersLogic(r.Context(), svcCtx)
		resp, err := l.GetAllUsers(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
