package Data

import (
	"github.com/team-zf/framework/utils"
)

// 洞府列表转为JsonMap输出格式
func (e *Person) __CaveToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range e.CaveList {
		result[utils.NewStringInt(k).ToString()] = v.ToJsonMap()
	}
	return result
}
