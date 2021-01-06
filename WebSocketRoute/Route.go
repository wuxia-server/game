package WebSocketRoute

import (
	"github.com/team-zf/framework/messages"
	"github.com/wuxia-server/game/WebSocketRoute/Client"
	"github.com/wuxia-server/game/WebSocketRoute/Cmd"
)

var (
	Route = messages.NewWebSocketMessageHandle()
)

func init() {
	Route.SetRoute(Cmd.Client_Connect, Client.M_Connect())
}
