package Dungeon

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/DataTable"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/Rule"
	"github.com/wuxia-server/game/StaticTable"
)

type Attack struct {
	Network.WebSocketRoute

	StoryId int // 关卡ID
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

	story := StaticTable.GetDungeonStory(e.StoryId)
	if story == nil {
		return Code.Dungeon_Attack_StoryNotExists
	}

	dungeon := person.GetDungeon(e.StoryId)
	// 没有权限攻打
	if dungeon == nil {
		return Code.Dungeon_Attack_NoRights
	}

	// ### 消耗体力
	{
		ddm, err := person.SubVigor(story.CostVigor)
		// 体力不足
		if err != nil {
			return Code.Dungeon_Attack_VigorInsufficient
		}
		e.Join(ddm)
	}

	// 固定通关, 通关星级随机
	star := utils.Range(1, 3)
	// 是否胜利
	win := star > 0
	// 是否为首次攻打
	first := dungeon.Star == 0

	if win {
		// ### 更新副本信息
		{
			if dungeon.Star < star {
				dungeon.Star = star
			}
			dungeon.AttackNum += 1
			dungeon.Save()
			e.Mod(Rule.RULE_DUNGEON, dungeon.ToJsonMap())
		}

		// ### 首次通关时生成下一关的数据
		if first {
			if next := StaticTable.GetDungeonStoryNext(e.StoryId); next != nil {
				dungeon := new(DataTable.UserDungeon)
				dungeon.Id = person.JoinToUserId(next.StoryId)
				dungeon.UserId = person.UserId()
				dungeon.StoryId = next.StoryId
				dungeon.Save()
				person.AddDungeon(dungeon)
				e.Mod(Rule.RULE_DUNGEON, dungeon.ToJsonMap())
			}
		}

		// ### 处理掉落
		{
			gainItems, ddm := person.Drop2(story.DropId)
			e.Join(ddm)
			e.Data("gain_items", gainItems)
		}

		// ### 获得经验
		{
			oldLevel := person.Level()
			ddm := person.AddExpV(story.Exp)
			e.Join(ddm)

			if person.Level() > oldLevel {
				// 升级了, 校验条件机制
			}
		}
	}

	return messages.RC_Success
}
