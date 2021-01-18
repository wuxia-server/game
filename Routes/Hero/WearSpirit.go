package Hero

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
)

type WearSpirit struct {
	Network.WebSocketRoute

	HeroId   int // 英雄ID
	SpiritId int // 玄气ID
	Slot     int // 槽位
}

func (e *WearSpirit) Parse() {
	e.HeroId = utils.NewStringAny(e.Params["hero_id"]).ToIntV()
	e.SpiritId = utils.NewStringAny(e.Params["spirit_id"]).ToIntV()
	e.Slot = utils.NewStringAny(e.Params["slot"]).ToIntV()
}

func (e *WearSpirit) Handle(agent *Network.WebSocketAgent) uint32 {
	// 无效的位置号
	if e.Slot < 1 || e.Slot > 4 {
		return Code.Hero_WearSpirit_SlotIncalid
	}

	return messages.RC_Success
}
