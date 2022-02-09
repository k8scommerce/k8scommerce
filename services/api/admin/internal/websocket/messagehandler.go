package websocket

// type MessageHandler struct {
// 	conn io.ReadWriteCloser
// }

// func (h *MessageHandler) Handle(conn){

// }

// type WsHandlerFunc func([]byte, ws.OpCode)

// func LoginWsHandler(svcCtx *svc.ServiceContext) WsHandlerFunc {
// 	return func(w ResponseWriter, r *Request) {
// 		var req types.MessageHandlerLoginRequest
// 		if err := httpx.Parse(r, &req); err != nil {
// 			httpx.Error(w, err)
// 			return
// 		}

// 		l := MessageHandlers.NewLoginLogic(r.Context(), svcCtx)
// 		resp, err := l.Login(req)
// 		if err != nil {

// 		} else {

// 		}
// 	}
// }
