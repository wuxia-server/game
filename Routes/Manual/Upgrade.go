package Manual

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Data"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/StaticTable"
)

type Upgrade struct {
	Network.WebSocketRoute

	ManualBankDetailId int // 图鉴库细节ID
}

func (e *Upgrade) Parse() {
	e.ManualBankDetailId = utils.NewStringAny(e.Params["manual_bank_detail_id"]).ToIntV()
}

func (e *Upgrade) Handle(agent *Network.WebSocketAgent) uint32 {
	person := Manage.GetPersonByAgent(agent)
	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	detail := StaticTable.GetManualBankDetail(e.ManualBankDetailId)
	// 没有这个图鉴细节
	if detail == nil {
		return Code.Manual_Upgrade_DetailNotExists
	}

	manual := person.GetManual(detail.Id)
	firstActive := manual == nil

	// ### 处理消耗
	{
		educate := StaticTable.GetManualEducate(1)
		if !firstActive {
			educate = StaticTable.GetManualEducate(manual.Level + 1)
		}
		// 图鉴等级已达上限
		if educate == nil {
			return Code.Manual_Upgrade_LevelUpperLimit
		}

		// 本次执行所需消耗的材料
		costs := make(map[int]int)
		// 魂魄ID, 必须
		costs[detail.SoulId] = educate.CostSoul
		// 额外消耗, 非必须
		if educate.ExtraCost != nil {
			costs[educate.ExtraCost.KeyToInt()] = educate.ExtraCost.ValueToInt()
		}

		ddm, err := person.SubItems(costs)
		if err != nil {
			return Code.Manual_Upgrade_UnderCost
		}
		e.Join(ddm)
	}

	// ### 处理升级
	{
		if firstActive {
			ddm, _ := person.AddManual(detail.Id)
			e.Join(ddm)
		} else {
			ddm, _ := person.UpgradeManual(detail.Id)
			e.Join(ddm)
		}
	}

	// ### 激活图鉴, 根据图鉴类型收获不同
	if firstActive {
		e.ActiveManual(detail, person)
	}

	return messages.RC_Success
}

func (e *Upgrade) ActiveManual(detail *StaticTable.ManualBankDetail, person *Data.Person) {
	switch detail.ActiveType {
	// 激活英雄
	case 1:
		{
			ddm, err := person.AddHero(detail.ActiveId)
			if err != nil {
				panic(err)
			}
			e.Join(ddm)
		}
	// 激活坐骑
	case 2:
		{
			// 待实现...
		}
	// 激活法宝
	case 3:
		{
			// 待实现...
		}
	}
}
