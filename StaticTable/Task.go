package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type Task struct {
	TaskId    int         `ST:"PK"`        // 任务组ID
	TaskType  int         `ST:"task_type"` // 任务组类型 (1、普通类型 2、每日刷新任务)
	DetailIds *Table.List `ST:"task_list"` // 细节ID列表
}

var (
	_TaskList []*Task
)

func init() {
	filePath := "./JSON/wx_task.json"
	rows, err := Table.LoadTable(filePath, &Task{})
	if err != nil {
		panic(err)
	}

	arr := make([]*Task, 0)
	for _, row := range rows {
		arr = append(arr, row.(*Task))
	}
	_TaskList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetTask(taskId int) (result *Task) {
	for _, row := range _TaskList {
		if row.TaskId == taskId {
			result = row
			break
		}
	}
	return
}

func GetTaskList() (result []*Task) {
	return _TaskList
}
