package Dungeon

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Data"
	"github.com/wuxia-server/game/DataTable"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/Rule"
	"github.com/wuxia-server/game/StaticTable"
)

type Attack struct {
	Network.WebSocketRoute

	StoryId int // 关卡ID

	_FirstAttack bool // 是否为首次攻打
	_BattleStar  int  // 战斗通关星级(0星代表未通关)
	_StStory     *StaticTable.DungeonStory
	_DtDungeon   *DataTable.UserDungeon
	_Person      *Data.Person
}

func (e *Attack) Parse() {
	e.StoryId = utils.NewStringAny(e.Params["story_id"]).ToIntV()
}

func (e *Attack) Handle(agent *Network.WebSocketAgent) uint32 {
	person := Manage.GetPersonByAgent(agent)
	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	stStory := StaticTable.GetDungeonStory(e.StoryId)
	if stStory == nil {
		return Code.Dungeon_Attack_StoryNotExists
	}

	dtDungeon := person.GetDungeon(e.StoryId)
	// 没有权限攻打
	if dtDungeon == nil {
		return Code.Dungeon_Attack_NoRights
	}

	// ### 消耗体力
	{
		ddm, err := person.SubVigor(stStory.CostVigor)
		// 体力不足
		if err != nil {
			return Code.Dungeon_Attack_VigorInsufficient
		}
		e.Join(ddm)
	}

	e._DtDungeon = dtDungeon
	e._StStory = stStory
	e._Person = person
	e._FirstAttack = dtDungeon.Star == 0
	e._BattleStar = e.Battle()

	e.UpdateDungeon()
	e.GenerateNextDungeon()
	e.Drop()
	e.GainExp()

	return messages.RC_Success
}

func (e *Attack) Battle() int {
	if true {
		// 固定通关, 通关星级随机
		return utils.Range(1, 3)
	} else {
		// 走正常Battle逻辑 (待处理)
		return 0
	}
}

// 更新副本信息
func (e *Attack) UpdateDungeon() {
	if e._BattleStar == 0 {
		return
	}
	if e._BattleStar > e._DtDungeon.Star {
		e._DtDungeon.Star = e._BattleStar
	}
	e._DtDungeon.AttackNum += 1
	e._DtDungeon.Save()
	e.Mod(Rule.RULE_DUNGEON, e._DtDungeon.ToJsonMap())
}

// 生成下一关
func (e *Attack) GenerateNextDungeon() {
	if e._BattleStar == 0 || e._FirstAttack {
		return
	}

	stNextStory := e._StStory.NextStory()
	if stNextStory != nil {
		dtDungeon := new(DataTable.UserDungeon)
		dtDungeon.Id = e._Person.JoinToUserId(stNextStory.StoryId)
		dtDungeon.UserId = e._Person.UserId()
		dtDungeon.StoryId = stNextStory.StoryId
		dtDungeon.Save()
		e._Person.AddDungeon(dtDungeon)
		e.Mod(Rule.RULE_DUNGEON, dtDungeon.ToJsonMap())
	}
}

// 掉落
func (e *Attack) Drop() {
	if e._BattleStar == 0 {
		return
	}
	gainItems, ddm := e._Person.Drop2(e._StStory.DropId)
	e.Join(ddm)
	e.Data("gain_items", gainItems)
}

// 获得经验
func (e *Attack) GainExp() {
	oldLevel := e._Person.Level()

	ddm := e._Person.AddExpV(e._StStory.Exp)
	e.Join(ddm)

	if e._Person.Level() > oldLevel {
		// 升级了, 校验条件机制
	}
}
