package Task

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
)

type Reward struct {
	Network.WebSocketRoute
}

func (e *Reward) Parse() {
}

func (e *Reward) Handle(agent *Network.WebSocketAgent) uint32 {
	return messages.RC_Success
}
