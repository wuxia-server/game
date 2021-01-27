package Data

import (
	"errors"
	"fmt"
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/DataTable"
	"github.com/wuxia-server/game/Rule"
)

// 任务列表转为JsonMap输出格式
func (e *Person) __TaskToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for _, task := range e.TaskList {
		result[utils.NewStringInt(task.DetailId).ToString()] = task.ToJsonMap()
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

// 添加或更新任务
func (e *Person) AddTask(detailId int) (*Network.WebSocketDDM, error) {
	task := e.GetTask(detailId)
	if task != nil {
		return nil, errors.New(fmt.Sprintf("用户(%d)已经拥有任务(%d), 无法重复获得.", e.UserId(), detailId))
	}

	task = DataTable.NewUserTask()
	task.Id = e.JoinToUserId(detailId)
	task.UserId = e.UserId()
	task.DetailId = detailId
	task.Status = 1
	task.Save()

	e.TaskList = append(e.TaskList, task)

	ddm := new(Network.WebSocketDDM)
	ddm.Mod(Rule.RULE_TASK, task.ToJsonMap())
	return ddm, nil
}
