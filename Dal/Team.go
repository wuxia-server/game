package Dal

import (
	"github.com/team-zf/framework/dal"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/DataTable"
)

func GetTeamListByUserId(userId int64) ([]*DataTable.UserTeam, error) {
	sqlstr := dal.MarshalGetSql(DataTable.NewUserTeam(), "user_id")
	rows, err := Control.GameDB.Query(sqlstr, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teamList := make([]*DataTable.UserTeam, 0)
	for rows.Next() {
		team := DataTable.NewUserTeam()
		err := rows.Scan(
			&team.Id,
			&team.UserId,
			&team.TeamId,
			&team.Position1,
			&team.Position2,
			&team.Position3,
			&team.Position4,
			&team.Position5,
			&team.DefendLundao,
			&team.DefendDongfu,
			&team.FightPower,
		)
		if err != nil {
			return nil, err
		}
		teamList = append(teamList, team)
	}
	return teamList, nil
}
