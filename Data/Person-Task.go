package Data

import (
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/DataTable"
)

// 任务列表转为JsonMap输出格式
func (e *Person) __TaskToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range e.TaskList {
		result[utils.NewStringInt(k).ToString()] = v.ToJsonMap()
	}
	return result
}

func (e *Person) GetTask(detailId int) (result *DataTable.UserTask) {
	for _, task := range e.TaskList {
		if task.DetailId == detailId {
			result = task
			break
		}
	}
	return
}
