package DataTable

import (
	"github.com/team-zf/framework/dal"
	"github.com/team-zf/framework/logger"
	"github.com/wuxia-server/game/Control"
	"time"
)

type UserCave struct {
	dal.BaseTable

	Id              int64     `db:"id,pk"`             // 主键 (用户ID+洞府ID)
	UserId          int64     `db:"user_id,!mod"`      // 用户ID
	CaveId          int       `db:"cave_id,!mod"`      // 洞府ID
	UpgradeId       int       `db:"upgrade_id"`        // 洞府等级ID
	Produce         int       `db:"produce"`           // 产量值
	UpdateTime      time.Time `db:"update_time"`       // 更新时间
	LastReceiveTime time.Time `db:"last_receive_time"` // 最后领取的时间
}

func (e *UserCave) GetTableName() (name string) {
	return "user_cave"
}

func (e *UserCave) ToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["cave_id"] = e.CaveId
	result["upgrade_id"] = e.UpgradeId
	result["produce"] = e.Produce
	result["update_time"] = e.UpdateTime.Unix()
	return result
}

func (e *UserCave) Save() {
	msg := &dal.DalMessage{
		UserId: e.UserId,
		Table:  e,
	}
	msg.RunFunc = func(db dal.IConnDB) error {
		_, err := db.Exec(dal.MarshalModSql(e),
			e.Id, e.UserId, e.CaveId,
			e.UpgradeId, e.Produce, e.UpdateTime, e.LastReceiveTime,
		)
		if err != nil {
			logger.Error("数据表[%s]保存失败, 错误原因: %+v", e.GetTableName(), err)
		}
		return err
	}
	Control.GameDB.AddMsg(msg)
}

func NewUserCave() *UserCave {
	result := new(UserCave)
	result.BaseTable.Init(result)
	return result
}
