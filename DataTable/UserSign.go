package DataTable

import (
	"github.com/team-zf/framework/dal"
	"github.com/team-zf/framework/logger"
	"github.com/wuxia-server/game/Control"
	"time"
)

type UserSign struct {
	dal.BaseTable

	Id       int64     `db:"id,pk"`        // 主键 (用户ID+天数)
	UserId   int64     `db:"user_id,!mod"` // 用户ID
	Day      int       `db:"day,!mod"`     // 月度自然天数
	Status   int       `db:"status"`       // 状态(0:未签 1:签到 2:补签)
	SignTime time.Time `db:"sign_time"`    // 签到时间
}

func (e *UserSign) GetTableName() (name string) {
	return "user_sign"
}

func (e *UserSign) ToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["day"] = e.Day
	result["status"] = e.Status
	result["sign_time"] = e.SignTime.Unix()
	return result
}

func (e *UserSign) Save() {
	msg := &dal.DalMessage{
		UserId: e.UserId,
		Table:  e,
	}
	msg.RunFunc = func(db dal.IConnDB) error {
		_, err := db.Exec(dal.MarshalModSql(e),
			e.Id, e.UserId, e.Day,
			e.Status, e.SignTime,
		)
		if err != nil {
			logger.Error("数据表[%s]保存失败, 错误原因: %+v", e.GetTableName(), err)
		}
		return err
	}
	Control.GameDB.AddMsg(msg)
}

func NewUserSign() *UserSign {
	result := new(UserSign)
	result.BaseTable.Init(result)
	return result
}
