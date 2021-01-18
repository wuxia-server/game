package DataTable

import (
	"github.com/team-zf/framework/dal"
	"github.com/team-zf/framework/logger"
	"github.com/wuxia-server/game/Control"
)

type UserDungeon struct {
	dal.BaseTable

	Id        int64 `db:"id,pk"`         // 主键 (用户ID+关卡ID)
	UserId    int64 `db:"user_id,!mod"`  // 用户ID
	StoryId   int   `db:"story_id,!mod"` // 关卡ID
	Star      int   `db:"star"`          // 最高通关星数
	AttackNum int   `db:"attack_num"`    // 已攻打次数
}

func (e *UserDungeon) GetTableName() (name string) {
	return "user_dungeon"
}

func (e *UserDungeon) ToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["story_id"] = e.StoryId
	result["star"] = e.Star
	result["attack_num"] = e.AttackNum
	return result
}

func (e *UserDungeon) Save() {
	msg := &dal.DalMessage{
		UserId: e.UserId,
		Table:  e,
	}
	msg.RunFunc = func(db dal.IConnDB) error {
		_, err := db.Exec(dal.MarshalModSql(e),
			e.Id, e.UserId, e.StoryId,
			e.Star, e.AttackNum,
		)
		if err != nil {
			logger.Error("数据表[%s]保存失败, 错误原因: %+v", e.GetTableName(), err)
		}
		return err
	}
	Control.GameDB.AddMsg(msg)
}

func NewUserDungeon() *UserDungeon {
	result := new(UserDungeon)
	result.BaseTable.Init(result)
	return result
}
