package Sign

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
)

type SupplementSign struct {
	Network.WebSocketRoute

	Day int // 补签日期(日)
}

func (e *SupplementSign) Parse() {
	e.Day = utils.NewStringAny(e.Params["day"]).ToIntV()
}

func (e *SupplementSign) Handle(agent *Network.WebSocketAgent) uint32 {
	return messages.RC_Success
}
