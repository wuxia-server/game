package DataTable

import (
	"github.com/team-zf/framework/dal"
	"github.com/team-zf/framework/logger"
	"github.com/wuxia-server/game/Control"
	"time"
)

type UserManual struct {
	dal.BaseTable

	Id         int64     `db:"id,pk"`          // 主键 (用户ID+图鉴ID)
	UserId     int64     `db:"user_id,!mod"`   // 用户ID
	ManualId   int       `db:"manual_id,!mod"` // 图鉴ID
	Level      int       `db:"level"`          // 等级
	UpdateTime time.Time `db:"update_time"`    // 更新时间
	ActiveTime time.Time `db:"active_time"`    // 激活时间
}

func (e *UserManual) GetTableName() (name string) {
	return "user_manual"
}

func (e *UserManual) ToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["manual_id"] = e.ManualId
	result["level"] = e.Level
	return result
}

func (e *UserManual) Save() {
	msg := &dal.DalMessage{
		UserId: e.UserId,
		Table:  e,
	}
	msg.RunFunc = func(db dal.IConnDB) error {
		_, err := db.Exec(dal.MarshalModSql(e),
			e.Id, e.UserId, e.ManualId,
			e.Level, e.UpdateTime, e.ActiveTime,
		)
		if err != nil {
			logger.Error("数据表[%s]保存失败, 错误原因: %+v", e.GetTableName(), err)
		}
		return err
	}
	Control.GameDB.AddMsg(msg)
}

func NewUserManual() *UserManual {
	result := new(UserManual)
	result.BaseTable.Init(result)
	return result
}
