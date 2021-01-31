package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
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
		if row.StoryId == storyId {
			result = row
			break
		}
	}
	return
}
