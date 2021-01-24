package Shop

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/Rule"
	"github.com/wuxia-server/game/StaticTable"
	"time"
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

	// ### 处理消耗
	{
		// 是否使用了刷新券
		ticketUsed := false

		// 校验刷新券的使用
		if stShop.RefreshTicketCost != nil {
			ddm, err := person.SubItem(stShop.RefreshTicketCost.KeyToInt(), stShop.RefreshTicketCost.ValueToInt())
			if err == nil {
				ticketUsed = true
				e.Join(ddm)
			}
		}

		// 校验普通消耗
		if stShop.RefreshItemCost != nil && !ticketUsed {
			ddm, err := person.SubItem(stShop.RefreshItemCost.KeyToInt(), stShop.RefreshItemCost.ValueToInt())
			if err != nil {
				return Code.Shop_Refresh_UnderCost
			}
			e.Join(ddm)
		}
	}

	// ### 处理刷新
	{
		dtShop.GenerateGoodsList()
		dtShop.UpdateTime = time.Now()
		dtShop.Save()
	}

	e.Mod(Rule.RULE_SHOP, dtShop.ToJsonMap())

	return messages.RC_Success
}
