package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/svc"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/zeromicro/go-zero/rest"
	"golang.org/x/net/context"
)

type WsHandleFunc func(*MessageHandler)

var handlers map[string]WsHandleFunc

func RegisterWsHandlers(svcCtx *svc.ServiceContext) map[string]WsHandleFunc {
	handlers = make(map[string]WsHandleFunc)
	handlers["/v1/api/ping"] = PingHandler(svcCtx)
	handlers["/v1/users/:currentPage/:pageSize"] = GetAllUsersHandler(svcCtx)

	return handlers
}

func RegisterWsUpgrade(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Locale},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/ws",
					Handler: WsUpgradeHandler(serverCtx),
				},
			}...,
		),
	)
}

func WsUpgradeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
			fmt.Println(err)
			return
		}

		RegisterWsHandlers(svcCtx)
		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					// handle error
				}
				if op.IsData() {
					m := NewMessageHandler(conn, msg, op)
					m.HandleMessage()

				}

				msg = nil
			}
		}()
	}
}

func NewMessageHandler(conn io.ReadWriteCloser, message []byte, op ws.OpCode) MessageHandler {
	return MessageHandler{conn: conn, op: op}
}

type MessageHandler struct {
	conn    io.ReadWriteCloser
	ctx     context.Context
	route   string
	message []byte
	op      ws.OpCode
	isErr   error
}

func (m *MessageHandler) Context() context.Context {
	if m.ctx != nil {
		return m.ctx
	}
	return context.Background()
}

func (m *MessageHandler) Error(err error) {

}

func (m *MessageHandler) HandleMessage() {
	var packet map[string]interface{}
	if err := json.Unmarshal(m.message, &packet); err != nil {
		m.isErr = err
		return
	}

	for key, payload := range packet {
		// payload
		if key == "path" {
			m.route = payload.(string)
			break
		}
	}

	if m.route == "" {
		return
	}

	// check for a matching payload type
	_, ok := handlers[m.route]
	if !ok {
		m.isErr = errors.New("unknown or invalid route")
		return
	}

	// call the handler since it exists
	handlers[m.route](m)
	fmt.Println("route:", m.route)
}

func (m *MessageHandler) toJSON(message interface{}) {
	if resp, err := json.Marshal(message); err == nil {
		fmt.Println(string(resp))
		err = wsutil.WriteServerMessage(m.conn, m.op, resp)
		if err != nil {
			// handle error
		}
	}
}
