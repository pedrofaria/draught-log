package handler

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/websocket"

	"github.com/pedrofaria/draught-log/internal/pkg/stream"
	"github.com/pedrofaria/draught-log/internal/pkg/stream/types"
)

var connectionUpgradeRegex = regexp.MustCompile(`(^|.*,\\s*)upgrade($|\\s*,)`)

func isWebsocketRequest(req *http.Request) bool {
	return connectionUpgradeRegex.MatchString(
		strings.ToLower(req.Header.Get("Connection")),
	) && strings.ToLower(req.Header.Get("Upgrade")) == "websocket"
}

type StreamHandler struct {
	manager *stream.Manager
}

func NewStreamHandler(m *stream.Manager) *StreamHandler {
	return &StreamHandler{
		manager: m,
	}
}

func (handler *StreamHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if isWebsocketRequest(r) {
		websocket.Handler(handler.ServeWebSocket).ServeHTTP(w, r)
		return
	}

	http.NotFound(w, r)
}

func (handler *StreamHandler) ServeWebSocket(ws *websocket.Conn) {
	ctx, cancel := context.WithCancel(ws.Request().Context())

	handler.manager.Stream(ctx, func(m types.Message) {
		if err := websocket.JSON.Send(ws, &m); err != nil {
			cancel()
			return
		}
	})
}
