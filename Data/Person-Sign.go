package Data

import (
	"github.com/team-zf/framework/utils"
)

// 签到列表转为JsonMap输出格式
func (e *Person) __SignToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range e.SignList {
		result[utils.NewStringInt(k).ToString()] = v.ToJsonMap()
	}
	return result
}
