package DataTable

import (
	"github.com/team-zf/framework/dal"
	"github.com/team-zf/framework/logger"
	"github.com/wuxia-server/game/Control"
	"time"
)

type UserItem struct {
	dal.BaseTable

	Id         int64     `db:"id,pk"`        // 主键自增列
	UserId     int64     `db:"user_id,!mod"` // 用户关联ID
	ItemId     int       `db:"item_id,!mod"` // 物品关联ID
	Num        int       `db:"num"`          // 数量
	UpdateTime time.Time `db:"update_time"`  // 更新时间
}

func (e *UserItem) GetTableName() (name string) {
	return "user_item"
}

func (e *UserItem) ToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["item_id"] = e.ItemId
	result["num"] = e.Num
	return result
}

func (e *UserItem) Save() {
	msg := &dal.DalMessage{
		UserId: e.UserId,
		Table:  e,
	}
	msg.RunFunc = func(db dal.IConnDB) error {
		_, err := db.Exec(dal.MarshalModSql(e),
			e.Id, e.UserId, e.ItemId,
			e.Num, e.UpdateTime,
		)
		if err != nil {
			logger.Error("数据表[%s]保存失败, 错误原因: %+v", e.GetTableName(), err)
		}
		return err
	}
	Control.GameDB.AddMsg(msg)
}

func NewUserItem() *UserItem {
	result := new(UserItem)
	result.BaseTable.Init(result)
	return result
}
