package Lottery

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
)

type Card struct {
	Network.WebSocketRoute

	LotteryId int // 抽奖表ID
}

func (e *Card) Parse() {
	e.LotteryId = utils.NewStringAny(e.Params["lottery_id"]).ToIntV()
}

func (e *Card) Handle(agent *Network.WebSocketAgent) uint32 {
	return messages.RC_Success
}
