package Shop

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/Rule"
	"github.com/wuxia-server/game/StaticTable"
	"math"
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
	person := Manage.GetPersonByAgent(agent)
	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	stShop := StaticTable.GetShop(e.ShopId)
	// 沒有这个商店
	if stShop == nil {
		return Code.Shop_Buy_ShopNotExists
	}

	dtShop := person.GetShop(e.ShopId)
	// 沒有这个商店
	if dtShop == nil {
		return Code.Shop_Buy_ShopNotExists
	}

	stDetail := StaticTable.GetShopDetail(e.DetailId)
	// 没有这个商品
	if stDetail == nil {
		return Code.Shop_Buy_DetailNotExists
	}

	dtDetail := dtShop.DetailList.GetDetailById(e.DetailId)
	// 没有这个商品
	if dtDetail == nil {
		return Code.Shop_Buy_DetailNotExists
	}

	// ### 处理购买条件
	{
		switch nc := stDetail.NormalCanbuy; nc {
		// 无限购买
		case -1:
		// 限量购买
		default:
			// 商品购买次数不足
			if dtDetail.Sales >= nc {
				return Code.Shop_Buy_UnderCanbuy
			}
		}

		// 购买和解锁条件暂未处理...
	}

	// ### 处理消耗
	{
		itemId := stDetail.GoodsCost.KeyToInt()
		itemNum := stDetail.GoodsCost.ValueToInt()

		if stDetail.DiscountCanbuy >= dtDetail.Sales {
			switch stDetail.Discount {
			// 售价打折
			case 1:
				n := stDetail.GoodsCost.ValueToFloat()
				p := utils.NewStringAny(stDetail.DiscountParams.At(0)).ToFloatV()
				itemNum = int(math.Ceil(n * p / 10000))
			// 全新售价打折(重新定义售价消耗ID和数量)
			case 2:
				cost := dtDetail.DiscountParams.(*Table.Map)
				itemId = cost.KeyToInt()
				itemNum = cost.ValueToInt()
			}
		}
		ddm, err := person.SubItem(itemId, itemNum)
		if err != nil {
			return Code.Shop_Buy_UnderCost
		}
		e.Join(ddm)
	}

	// ### 处理购买
	{
		dtDetail.Sales += 1
		dtShop.Save()
		e.Mod(Rule.RULE_SHOP, dtShop.ToJsonMap())
	}

	// ### 处理商品到账
	{
		gainItems, ddm := person.Drop2(stDetail.GoodsDropId)
		e.Join(ddm)
		e.Data("gain_items", gainItems)
	}

	return messages.RC_Success
}
