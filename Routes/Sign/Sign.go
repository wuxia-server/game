package Sign

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
)

type Sign struct {
	Network.WebSocketRoute
}

func (e *Sign) Parse() {
}

func (e *Sign) Handle(agent *Network.WebSocketAgent) uint32 {
	return messages.RC_Success
}
