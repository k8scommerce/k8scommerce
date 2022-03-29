package websocket

import (
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/logic/api"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/logic/users"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/types"
)

func GetAllUsersHandler(svcCtx *svc.ServiceContext) WsHandleFunc {
	return func(m *MessageHandler) {
		var req types.GetAllUsersRequest
		if err := Parse(m, &req); err != nil {
			// wsx.Error(w, err)
			return
		}

		l := users.NewGetAllUsersLogic(m.Context(), svcCtx)
		resp, err := l.GetAllUsers(req)
		if err != nil {
			m.Error(err)
		} else {
			m.toJSON(resp)
		}
	}
}

func PingHandler(svcCtx *svc.ServiceContext) WsHandleFunc {
	return func(m *MessageHandler) {
		l := api.NewPingLogic(m.Context(), svcCtx)
		resp, err := l.Ping()
		if err != nil {
			m.Error(err)
		} else {
			m.toJSON(resp)
		}
	}
}
