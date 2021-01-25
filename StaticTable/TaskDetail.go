package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type TaskDetail struct {
	DetailId     int `ST:"PK"`        // 细节ID
	TaskCondId   int `ST:"task_cond"` // 任务完成条件ID
	TaskDropId   int `ST:"task_drop"` // 任务掉落ID
	NextDetailId int `ST:"next_task"` // 下一个细节ID
}

var (
	_TaskDetailList []*TaskDetail
)

func init() {
	filePath := "./JSON/wx_task_detail.json"
	rows, err := Table.LoadTable(filePath, &TaskDetail{})
	if err != nil {
		panic(err)
	}

	arr := make([]*TaskDetail, 0)
	for _, row := range rows {
		arr = append(arr, row.(*TaskDetail))
	}
	_TaskDetailList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetTaskDetail(detailId int) (result *TaskDetail) {
	for _, row := range _TaskDetailList {
		if row.DetailId == detailId {
			result = row
			break
		}
	}
	return
}
