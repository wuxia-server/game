package Hero

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Manage"
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
	dtItem := person.GetItem(e.ItemId)
	// 没有这个物品
	if stItem == nil || dtItem == nil {
		return Code.Hero_FeedExp_ItemNotExists
	}

	stHero := StaticTable.GetHero(e.HeroId)
	dtHero := person.GetHero(e.HeroId)
	// 没有这个英雄
	if stHero == nil || dtHero == nil {
		return Code.Hero_FeedExp_HeroNotExists
	}

	// 数量不足
	if dtItem.Num == 0 {
		return Code.Hero_FeedExp_QuantityInsufficient
	}

	// 该英雄经验已满
	if dtHero.Level >= 99 {
		return Code.Hero_FeedExp_HeroExpSpiledOver
	}

	return messages.RC_Success
}
