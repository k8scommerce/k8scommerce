package websocket

import (
	"k8scommerce/services/api/admin/internal/logic/api"
	"k8scommerce/services/api/admin/internal/logic/users"
	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
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
