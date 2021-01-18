package Data

import (
	"github.com/team-zf/framework/utils"
)

// 图鉴列表转为JsonMap输出格式
func (e *Person) __ManualToJson() map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range e.ManualList {
		result[utils.NewStringInt(k).ToString()] = v.Level
	}
	return result
}
