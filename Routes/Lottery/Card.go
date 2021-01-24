package Lottery

import (
	"errors"
	"fmt"
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/StaticTable"
)

type Card struct {
	Network.WebSocketRoute

	LotteryId int // 抽奖表ID

	_StLottery *StaticTable.Lottery
}

func (e *Card) Parse() {
	e.LotteryId = utils.NewStringAny(e.Params["lottery_id"]).ToIntV()
}

func (e *Card) Handle(agent *Network.WebSocketAgent) uint32 {
	person := Manage.GetPersonByAgent(agent)
	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	stLottery := StaticTable.GetLottery(e.LotteryId)
	// 没有这个抽奖项
	if stLottery == nil {
		return Code.Lottery_Card_LotteryNotExists
	}

	// ### 处理消耗
	{
		ddm, err := person.SubItem(
			stLottery.LotteryCost.KeyToInt(),
			stLottery.LotteryCost.ValueToInt(),
		)
		// 抽奖成本不足
		if err != nil {
			return Code.Lottery_Card_UnderCost
		}
		e.Join(ddm)
	}

	// ### 处理抽卡
	{
		gainItems := make([][]int, 0)
		details := StaticTable.GetLotteryDetailList(stLottery.LotteryBank)
		for i := 0; i < stLottery.LotteryCount; i++ {
			np := 0
			prob := utils.PercentV()

			var detail *StaticTable.LotteryDetail
			for _, v := range details {
				if v.DropProb+np >= prob {
					detail = v
					break
				}
				np += v.DropProb
			}

			if detail == nil {
				panic(errors.New(fmt.Sprintf("抽奖库(%d)的所有细节项的掉落概率总和低于100%", stLottery.LotteryBank)))
			}

			dropItems, ddm := person.Drop(detail.DropId)
			for _, dropItem := range dropItems {
				gainItems = append(gainItems, dropItem.ToArray())
			}
			e.Join(ddm)
		}
		e.Data("gain_items", gainItems)
	}

	return messages.RC_Success
}
