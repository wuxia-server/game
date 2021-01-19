package Cave

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Const"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/Rule"
	"github.com/wuxia-server/game/StaticTable"
	"math"
	"time"
)

type Upgrade struct {
	Network.WebSocketRoute

	CaveId int // 洞府ID
}

func (e *Upgrade) Parse() {
	e.CaveId = utils.NewStringAny(e.Params["cave_id"]).ToIntV()
}

func (e *Upgrade) Handle(agent *Network.WebSocketAgent) uint32 {
	person := Manage.GetPersonByAgent(agent)
	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	stCave := StaticTable.GetCave(e.CaveId)
	// 没有这个洞府
	if stCave == nil {
		return Code.Cave_Upgrade_CaveNotExists
	}

	dtCave := person.GetCave(e.CaveId)
	// 没有这个洞府
	if dtCave == nil {
		return Code.Cave_Upgrade_CaveNotExists
	}

	stCaveLv := StaticTable.GetCaveLv(dtCave.UpgradeId)
	// 无法升级, 已满级或达到了角色等级
	if stCaveLv.NextId == -1 || stCaveLv.Level >= person.User.Level {
		return Code.Cave_Upgrade_AlreadyFullLevel
	}

	nowtime := time.Now()

	// ### 处理消耗
	{
		costs := make(map[int]int)
		arr, _ := stCaveLv.Cost.ToMapArray()
		for _, m := range arr {
			costs[m.KeyToInt()] = m.ValueToInt()
		}
		ddm, err := person.SubItems(costs)
		if err != nil {
			return Code.Cave_Upgrade_UnderCost
		}
		e.Join(ddm)
	}

	// ### 更新洞府产量值
	{
		// 相差纳秒值
		diffNano := nowtime.Sub(dtCave.UpdateTime)
		// 相差秒值
		diffSecond := int(math.Floor(float64(diffNano) / float64(time.Second)))

		// 最大产量值
		maxProduce := stCaveLv.Maximum * Const.CaveProduceEGM
		// 每秒产量值
		perProduce := stCaveLv.Speed * Const.CaveProduceEGM / 60 / 60
		// 新的产量值
		newProduce := dtCave.Produce + diffSecond*perProduce

		if newProduce >= maxProduce {
			dtCave.Produce = maxProduce
			dtCave.UpdateTime = nowtime
		} else {
			dtCave.Produce = newProduce
			dtCave.UpdateTime = dtCave.UpdateTime.Add(time.Duration(diffSecond) * time.Second)
		}
	}

	// ### 处理升级
	{
		dtCave.UpgradeId = stCaveLv.NextId
		dtCave.Save()
		e.Mod(Rule.RULE_CAVE, dtCave.ToJsonMap())
	}

	return messages.RC_Success
}
