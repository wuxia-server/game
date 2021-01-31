package Dungeon

import (
	"encoding/json"
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Data"
	"github.com/wuxia-server/game/DataTable"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/Rule"
	"github.com/wuxia-server/game/StaticTable"
	"io/ioutil"
	"os"
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

	// 挑战次数不足
	if dungeon.AttackNum >= story.AttackNum && story.AttackNum > -1 {
		return Code.Dungeon_Attack_AttackNumInsufficient
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

	// 固定胜利, 胜利星级随机
	star := utils.Range(1, 3)
	{
		datas := e.Battle()
		for k, v := range datas {
			e.Data(k, v)
		}
		result := datas["result"].(map[string]interface{})
		star = utils.NewStringAny(result["star"]).ToIntV()
	}

	// 是否胜利
	win := star > 0

	if win {
		// 首次胜利
		if dungeon.Star == 0 {
			switch {
			// 通关章节
			case story.NextId == -1:
				switch story.Type {
				// 普通章节通关, 开启下一章
				case 1:
					e.OpenNextChapter(story, person)
					e.OpenEliteStory(story, person)
					// 做条件校验
					e.Join(person.CondVerify())
				// 精英章节通关
				case 2:
				}
			// 普通小关卡过关, 生成下一关卡
			default:
				e.OpenNextStory(story, person)
			}
		}

		// ### 更新副本信息
		{
			if dungeon.Star < star {
				dungeon.Star = star
			}
			dungeon.AttackNum += 1
			dungeon.Save()
			e.Mod(Rule.RULE_DUNGEON, dungeon.ToJsonMap())
		}

		// ### 处理掉落
		{
			gainItems, ddm := person.Drop2(story.DropId)
			e.Join(ddm)
			e.Data("gain_items", gainItems)
		}

		// ### 获得经验
		{
			ddm := person.AddExpV(story.Exp)
			e.Join(ddm)
		}
	}

	return messages.RC_Success
}

// 开启下一章
func (e *Attack) OpenNextChapter(story *StaticTable.DungeonStory, person *Data.Person) {
	chapter := StaticTable.GetDungeonChapter(story.ChapterId)
	next := StaticTable.GetDungeonChapter(chapter.NextChapterId)
	if next != nil {
		story := StaticTable.GetDungeonStory(next.NormalStoryId)
		dungeon := DataTable.NewUserDungeon()
		dungeon.Id = person.JoinToUserId(story.StoryId)
		dungeon.UserId = person.UserId()
		dungeon.StoryId = story.StoryId
		dungeon.Save()
		person.AddDungeon(dungeon)
		e.Mod(Rule.RULE_DUNGEON, dungeon.ToJsonMap())
	}
}

// 开启精英关卡
func (e *Attack) OpenEliteStory(story *StaticTable.DungeonStory, person *Data.Person) {
	chapter := StaticTable.GetDungeonChapter(story.ChapterId)
	if chapter.EliteStoryId != -1 {
		story := StaticTable.GetDungeonStory(chapter.EliteStoryId)
		dungeon := DataTable.NewUserDungeon()
		dungeon.Id = person.JoinToUserId(story.StoryId)
		dungeon.UserId = person.UserId()
		dungeon.StoryId = story.StoryId
		dungeon.Save()
		person.AddDungeon(dungeon)
		e.Mod(Rule.RULE_DUNGEON, dungeon.ToJsonMap())
	}
}

// 开启下一关卡
func (e *Attack) OpenNextStory(story *StaticTable.DungeonStory, person *Data.Person) {
	next := StaticTable.GetDungeonStory(story.NextId)
	dungeon := DataTable.NewUserDungeon()
	dungeon.Id = person.JoinToUserId(next.StoryId)
	dungeon.UserId = person.UserId()
	dungeon.StoryId = next.StoryId
	dungeon.Save()
	person.AddDungeon(dungeon)
	e.Mod(Rule.RULE_DUNGEON, dungeon.ToJsonMap())
}

func (e *Attack) Battle() map[string]interface{} {
	file, err := os.Open("./battle_test.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buff, _ := ioutil.ReadAll(file)
	datas := make(map[string]interface{})
	json.Unmarshal(buff, &datas)
	return datas
}
