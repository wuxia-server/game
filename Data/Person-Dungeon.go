package Data

import (
	"github.com/team-zf/framework/utils"
)

// 副本列表转为JsonMap输出格式
func (e *Person) __DungeonToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range e.DungeonList {
		result[utils.NewStringInt(k).ToString()] = v.ToJsonMap()
	}
	return result
}
