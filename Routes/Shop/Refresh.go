package Shop

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/StaticTable"
)

type Refresh struct {
	Network.WebSocketRoute

	ShopId int // 商店ID
}

func (e *Refresh) Parse() {
	e.ShopId = utils.NewStringAny(e.Params["shop_id"]).ToIntV()
}

func (e *Refresh) Handle(agent *Network.WebSocketAgent) uint32 {
	person := Manage.GetPersonByAgent(agent)
	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	stShop := StaticTable.GetShop(e.ShopId)
	// 沒有这个商店
	if stShop == nil {
		return Code.Shop_Refresh_ShopNotExists
	}

	dtShop := person.GetShop(e.ShopId)
	// 沒有开通这个商店
	if dtShop == nil {
		return Code.Shop_Refresh_ShopNotOpen
	}

	// 该商店不支持手动刷新
	if !stShop.SupportRefresh {
		return Code.Shop_Refresh_NotSupportRefresh
	}

	return messages.RC_Success
}
