package Dal

import (
	"github.com/team-zf/framework/dal"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/DataTable"
)

func GetTaskListByUserId(userId int64) ([]*DataTable.UserTask, error) {
	sqlstr := dal.MarshalGetSql(DataTable.NewUserTask(), "user_id")
	rows, err := Control.GameDB.Query(sqlstr, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	taskList := make([]*DataTable.UserTask, 0)
	for rows.Next() {
		task := DataTable.NewUserTask()
		err := rows.Scan(
			&task.Id,
			&task.UserId,
			&task.DetailId,
			&task.Status,
		)
		if err != nil {
			return nil, err
		}
		taskList = append(taskList, task)
	}
	return taskList, nil
}
