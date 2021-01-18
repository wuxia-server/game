package Shop

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
)

type Refresh struct {
	Network.WebSocketRoute

	ShopId int // 商店ID
}

func (e *Refresh) Parse() {
	e.ShopId = utils.NewStringAny(e.Params["shop_id"]).ToIntV()
}

func (e *Refresh) Handle(agent *Network.WebSocketAgent) uint32 {
	return messages.RC_Success
}
