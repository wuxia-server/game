package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/utils"
)

type DungeonStory struct {
	StoryId     int         `ST:"PK"`          // 关卡ID
	ChapterId   int         `ST:"chapter"`     // 所属章节ID
	NextId      int         `ST:"next"`        // 下一关卡ID
	Type        int         `ST:"type"`        // 关卡类型
	AttackNum   int         `ST:"attack_num"`  // 可挑战次数
	CostVigor   int         `ST:"cost_energy"` // 消耗体力值
	Exp         int         `ST:"exp"`         // 奖励经验
	DropId      int         `ST:"drop_id"`     // 掉落ID
	MonsterList *Table.List `ST:"enemy_list"`  // 怪物列表
}

func (e *DungeonStory) NextStory() *DungeonStory {
	if e.NextId == -1 {
		nextChapterId := GetDungeonChapter(e.ChapterId).NextChapterId
		if nextChapterId == -1 {
			return nil
		} else {
			if e.Type == 1 {
				return GetDungeonStory(GetDungeonChapter(nextChapterId).NormalStoryId)
			} else {
				return GetDungeonStory(GetDungeonChapter(nextChapterId).EliteStoryId)
			}
		}
	} else {
		return GetDungeonStory(e.NextId)
	}
}

var (
	_DungeonStoryList []*DungeonStory
)

func init() {
	filePath := "./JSON/wx_dungeon_story.json"
	rows, err := Table.LoadTable(filePath, &DungeonStory{})
	if err != nil {
		panic(err)
	}

	arr := make([]*DungeonStory, 0)
	for _, row := range rows {
		arr = append(arr, row.(*DungeonStory))
	}
	_DungeonStoryList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetDungeonStory(storyId int) (result *DungeonStory) {
	for _, row := range _DungeonStoryList {
		if row.ChapterId == storyId {
			newrow := utils.ReflectNew(row)
			result = newrow.(*DungeonStory)
			break
		}
	}
	return
}