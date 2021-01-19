package Cave

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Const"
	"github.com/wuxia-server/game/Data"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/Rule"
	"github.com/wuxia-server/game/StaticTable"
	"math"
	"time"
)

type Harvest struct {
	Network.WebSocketRoute

	CaveId int // 洞府ID
}

func (e *Harvest) Parse() {
	e.CaveId = utils.NewStringAny(e.Params["cave_id"]).ToIntV()
}

func (e *Harvest) Handle(agent *Network.WebSocketAgent) uint32 {
	person := Manage.GetPersonByAgent(agent)
	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	stCave := StaticTable.GetCave(e.CaveId)
	// 没有这个洞府
	if stCave == nil {
		return Code.Cave_Harvest_CaveNotExists
	}

	dtCave := person.GetCave(e.CaveId)
	// 没有这个洞府
	if dtCave == nil {
		return Code.Cave_Harvest_CaveNotExists
	}

	nowtime := time.Now()
	// 丰收产量值
	produce := 0

	// ### 防止频繁丰收
	{
		// 相差纳秒值
		diffNano := nowtime.Sub(dtCave.LastReceiveTime)
		// 相差秒值
		diffSecond := int(math.Floor(float64(diffNano) / float64(time.Second)))

		// 没到丰收时间
		if diffSecond < stCave.RewardTime {
			return Code.Cave_Harvest_NotYetHarvestTime
		}
	}

	// ### 更新洞府产量值
	{
		// 相差纳秒数
		diffNano := nowtime.Sub(dtCave.UpdateTime)
		// 相差秒数
		diffSecond := int(math.Floor(float64(diffNano) / float64(time.Second)))

		stCaveLv := StaticTable.GetCaveLv(stCave.UpgradeId)

		// 最大产量值
		maxProduce := stCaveLv.Maximum * Const.CaveProduceEGM
		// 每秒产量值
		perProduce := stCaveLv.Speed * Const.CaveProduceEGM / 60 / 60
		// 新的产量值
		newProduce := dtCave.Produce + (diffSecond * perProduce)

		if newProduce >= maxProduce {
			dtCave.Produce = maxProduce
			dtCave.UpdateTime = nowtime
		} else {
			dtCave.Produce = newProduce
			dtCave.UpdateTime = dtCave.UpdateTime.Add(time.Second * time.Duration(diffSecond))
		}

		// 转为正常倍数(向下取整)
		produce = int(math.Floor(float64(dtCave.Produce) / float64(Const.CaveProduceEGM)))

		dtCave.Produce = produce
		dtCave.LastReceiveTime = nowtime
		dtCave.Save()

		e.Mod(Rule.RULE_CAVE, dtCave.ToJsonMap())
	}

	// ### 处理掉落
	{
		dropItem := new(Data.DropItem)
		dropItem.ItemId = stCave.RewardItemId
		dropItem.Num = produce

		// 获得物品列表
		e.Data("gain_items", [][]int{dropItem.ToArray()})

		e.Join(person.AddItemV(dropItem.ItemId, dropItem.Num))
	}

	return messages.RC_Success
}
