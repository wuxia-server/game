package DataTable

import (
	"github.com/team-zf/framework/dal"
	"github.com/team-zf/framework/logger"
	"github.com/wuxia-server/game/Control"
)

type UserTeam struct {
	dal.BaseTable

	Id           int64 `db:"id,pk"`         // 主键 (用户ID+队伍ID)
	UserId       int64 `db:"user_id,!mod"`  // 用户ID
	TeamId       int   `db:"team_id,!mod"`  // 队伍ID
	Position1    int   `db:"pos_1"`         // 队伍位置1英雄ID
	Position2    int   `db:"pos_2"`         // 队伍位置2英雄ID
	Position3    int   `db:"pos_3"`         // 队伍位置3英雄ID
	Position4    int   `db:"pos_4"`         // 队伍位置4英雄ID
	Position5    int   `db:"pos_5"`         // 队伍位置5英雄ID
	DefendLundao int   `db:"defend_lundao"` // 是否为守护论道的队伍
	DefendDongfu int   `db:"defend_dongfu"` // 是否为守护洞府的队伍
	FightPower   int   `db:"fight_power"`   // 总战斗力
}

func (e *UserTeam) GetTableName() (name string) {
	return "user_team"
}

func (e *UserTeam) ToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["team_id"] = e.TeamId
	result["pos_1"] = e.Position1
	result["pos_2"] = e.Position2
	result["pos_3"] = e.Position3
	result["pos_4"] = e.Position4
	result["pos_5"] = e.Position5
	result["defend_lundao"] = e.DefendLundao
	result["defend_dongfu"] = e.DefendDongfu
	result["fight_power"] = e.FightPower
	return result
}

func (e *UserTeam) Save() {
	msg := &dal.DalMessage{
		UserId: e.UserId,
		Table:  e,
	}
	msg.RunFunc = func(db dal.IConnDB) error {
		_, err := db.Exec(dal.MarshalModSql(e),
			e.Id, e.UserId, e.TeamId,
			e.Position1, e.Position2, e.Position3, e.Position4, e.Position5,
			e.DefendLundao, e.DefendDongfu, e.FightPower,
		)
		if err != nil {
			logger.Error("数据表[%s]保存失败, 错误原因: %+v", e.GetTableName(), err)
		}
		return err
	}
	Control.GameDB.AddMsg(msg)
}

func NewUserTeam() *UserTeam {
	result := new(UserTeam)
	result.BaseTable.Init(result)
	return result
}
