package Dal

import (
	"github.com/team-zf/framework/dal"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/DataTable"
)

func GetDungeonListByUserId(userId int64) ([]*DataTable.UserDungeon, error) {
	sqlstr := dal.MarshalGetSql(DataTable.NewUserDungeon(), "user_id")
	rows, err := Control.GameDB.Query(sqlstr, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dungeonList := make([]*DataTable.UserDungeon, 0)
	for rows.Next() {
		dungeon := DataTable.NewUserDungeon()
		err := rows.Scan(
			&dungeon.Id,
			&dungeon.UserId,
			&dungeon.StoryId,
			&dungeon.Star,
			&dungeon.AttackNum,
		)
		if err != nil {
			return nil, err
		}
		dungeonList = append(dungeonList, dungeon)
	}
	return dungeonList, nil
}
