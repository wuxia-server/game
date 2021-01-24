package DataTable

import (
	"github.com/team-zf/framework/dal"
	"github.com/team-zf/framework/logger"
	"github.com/wuxia-server/game/Control"
)

type UserTask struct {
	dal.BaseTable

	Id       int64 `db:"id,pk"`          // 主键 (用户ID+任务细节ID)
	UserId   int64 `db:"user_id,!mod"`   // 用户ID
	DetailId int   `db:"detail_id,!mod"` // 任务细节ID
	Status   int   `db:"status"`         // 状态(1:达成条件 2:已领取)
}

func (e *UserTask) GetTableName() (name string) {
	return "user_task"
}

func (e *UserTask) ToJsonMap() map[string]interface{} {
	return map[string]interface{}{
		"task_id": e.DetailId,
		"status":  e.Status,
	}
}

func (e *UserTask) Save() {
	msg := &dal.DalMessage{
		UserId: e.UserId,
		Table:  e,
	}
	msg.RunFunc = func(db dal.IConnDB) error {
		_, err := db.Exec(dal.MarshalModSql(e),
			e.Id, e.UserId, e.DetailId,
			e.Status,
		)
		if err != nil {
			logger.Error("数据表[%s]保存失败, 错误原因: %+v", e.GetTableName(), err)
		}
		return err
	}
	Control.GameDB.AddMsg(msg)
}

func NewUserTask() *UserTask {
	result := new(UserTask)
	result.BaseTable.Init(result)
	return result
}
