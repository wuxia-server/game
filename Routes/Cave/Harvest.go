package Cave

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
)

type Harvest struct {
	Network.WebSocketRoute

	CaveId int // 洞府ID
}

func (e *Harvest) Parse() {
	e.CaveId = utils.NewStringAny(e.Params["cave_id"]).ToIntV()
}

func (e *Harvest) Handle(agent *Network.WebSocketAgent) uint32 {
	return messages.RC_Success
}
