package Data

import (
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/DataTable"
)

// 洞府列表转为JsonMap输出格式
func (e *Person) __CaveToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for _, cave := range e.CaveList {
		result[utils.NewStringInt(cave.CaveId).ToString()] = cave.ToJsonMap()
	}
	return result
}

func (e *Person) GetCave(caveId int) (result *DataTable.UserCave) {
	for _, cave := range e.CaveList {
		if cave.CaveId == caveId {
			result = cave
			break
		}
	}
	return
}
