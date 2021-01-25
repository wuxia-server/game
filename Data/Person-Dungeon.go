package Data

import (
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/DataTable"
)

// 副本列表转为JsonMap输出格式
func (e *Person) __DungeonToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for _, dungeon := range e.DungeonList {
		result[utils.NewStringInt(dungeon.StoryId).ToString()] = dungeon.ToJsonMap()
	}
	return result
}

func (e *Person) GetDungeon(storyId int) (result *DataTable.UserDungeon) {
	for _, dungeon := range e.DungeonList {
		if dungeon.StoryId == storyId {
			result = dungeon
			break
		}
	}
	return
}

func (e *Person) AddDungeon(dungeon *DataTable.UserDungeon) {
	if e.DungeonList == nil {
		e.DungeonList = make([]*DataTable.UserDungeon, 0)
	}
	e.DungeonList = append(e.DungeonList, dungeon)
}
