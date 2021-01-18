package Shop

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
)

type Buy struct {
	Network.WebSocketRoute

	ShopId   int // 商店ID
	DetailId int // 商店细节ID(商品ID)
}

func (e *Buy) Parse() {
	e.ShopId = utils.NewStringAny(e.Params["shop_id"]).ToIntV()
	e.DetailId = utils.NewStringAny(e.Params["detail_id"]).ToIntV()
}

func (e *Buy) Handle(agent *Network.WebSocketAgent) uint32 {
	return messages.RC_Success
}
