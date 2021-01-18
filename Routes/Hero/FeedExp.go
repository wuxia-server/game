package Hero

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/tables"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Manage"
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
	client := Manage.GetPersonByAgent(agent)

	tItem := tables.GetTable("item").ByKey(e.ItemId)
	tHero := tables.GetTable("hero").ByKey(e.HeroId)
	item := client.GetItemById(e.ItemId)
	hero := client.GetHeroById(e.HeroId)

	// 没有这个物品
	if tItem == nil || item == nil {
		return Code.Hero_FeedExp_ItemNotExists
	}

	// 没有这个英雄
	if tHero == nil || hero == nil {
		return Code.Hero_FeedExp_HeroNotExists
	}

	// 数量不足
	if item.Num == 0 {
		return Code.Hero_FeedExp_QuantityInsufficient
	}

	// 该英雄经验已满
	if hero.Level >= 99 {
		return Code.Hero_FeedExp_HeroExpSpiledOver
	}

	return messages.RC_Success
}
