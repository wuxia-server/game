package Hero

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Const"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/Rule"
	"github.com/wuxia-server/game/StaticTable"
)

type FeedExp struct {
	Network.WebSocketRoute

	HeroId int // 英雄ID
	ItemId int // 物品ID
}

func (e *FeedExp) Parse() {
	e.HeroId = utils.NewStringAny(e.Params["hero_id"]).ToIntV()
	e.ItemId = utils.NewStringAny(e.Params["item_id"]).ToIntV()
}

func (e *FeedExp) Handle(agent *Network.WebSocketAgent) uint32 {
	person := Manage.GetPersonByAgent(agent)
	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	stItem := StaticTable.GetItem(e.ItemId)
	// 没有这个物品
	if stItem == nil {
		return Code.Hero_FeedExp_ItemNotExists
	}

	// 无效的物品
	if stItem.ItemType != Const.ItemType_ExpCard {
		return Code.Hero_FeedExp_InvalidItem
	}

	stHero := StaticTable.GetHero(e.HeroId)
	dtHero := person.GetHero(e.HeroId)
	// 没有这个英雄
	if stHero == nil || dtHero == nil {
		return Code.Hero_FeedExp_HeroNotExists
	}

	// 数量不足
	dtItem := person.GetItem(e.ItemId)
	if dtItem == nil || dtItem.Num == 0 {
		return Code.Hero_FeedExp_QuantityInsufficient
	}

	hlv := StaticTable.GetHeroLvByExp(dtHero.Exp)
	nlv := StaticTable.GetHeroLv(hlv.Level + 1)
	// 这个英雄的经验已经满了
	if dtHero.Exp == nlv.NeedExp-1 {
		return Code.Hero_FeedExp_HeroExpFull
	}

	// ### 处理消耗
	{
		ddm, err := person.SubItem(e.ItemId, 1)
		if err != nil {
			return Code.Hero_FeedExp_QuantityInsufficient
		}
		e.Join(ddm)
	}

	// ### 处理增加经验
	{
		// 原始等级
		originLevel := dtHero.Level

		exp := dtHero.Exp + stItem.RelatedValue
		hlv := StaticTable.GetHeroLvByExp(exp)
		// 超出角色等级
		if hlv.Level > person.Level() {
			hlv := StaticTable.GetHeroLv(person.Level())
			nlv := StaticTable.GetHeroLv(person.Level() + 1)
			dtHero.Level = hlv.Level
			dtHero.Exp = nlv.NeedExp - 1
		} else {
			dtHero.Level = hlv.Level
			dtHero.Exp = exp
		}

		// 升级了
		if originLevel < dtHero.Level {
			// 更新属性(待定)
		}

		dtHero.Save()

		e.Mod(Rule.RULE_HERO, dtHero.ToJsonMap())

		// 告知客户端本次升了几级
		e.Data("level_up", dtHero.Level-originLevel)
	}

	return messages.RC_Success
}
