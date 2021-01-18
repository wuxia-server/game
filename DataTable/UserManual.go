package DataTable

import (
	"github.com/team-zf/framework/dal"
	"time"
)

type UserManual struct {
	dal.BaseTable

	Id         int64     `db:"id,pk"`          // 主键 (用户ID+图鉴ID)
	UserId     int       `db:"user_id,!mod"`   // 用户ID
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

func NewUserManual() *UserManual {
	result := new(UserManual)
	result.BaseTable.Init(result)
	return result
}
