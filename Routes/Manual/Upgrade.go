package Manual

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
)

type Upgrade struct {
	Network.WebSocketRoute

	ManualBankDetailId int // 图鉴库细节ID
}

func (e *Upgrade) Parse() {
	e.ManualBankDetailId = utils.NewStringAny(e.Params["manual_bank_detail_id"]).ToIntV()
}

func (e *Upgrade) Handle(agent *Network.WebSocketAgent) uint32 {
	return messages.RC_Success
}
