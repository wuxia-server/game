package Dal

import (
	"github.com/team-zf/framework/dal"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/DataTable"
)

func GetCaveListByUserId(userId int64) ([]*DataTable.UserCave, error) {
	sqlstr := dal.MarshalGetSql(DataTable.NewUserCave(), "user_id")
	rows, err := Control.GameDB.Query(sqlstr, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	caveList := make([]*DataTable.UserCave, 0)
	for rows.Next() {
		cave := DataTable.NewUserCave()
		err := rows.Scan(
			&cave.Id,
			&cave.UserId,
			&cave.CaveId,
			&cave.UpgradeId,
			&cave.Produce,
			&cave.UpdateTime,
			&cave.LastReceiveTime,
		)
		if err != nil {
			return nil, err
		}
		caveList = append(caveList, cave)
	}
	return caveList, nil
}
