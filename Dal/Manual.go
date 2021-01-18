package Dal

import (
	"github.com/team-zf/framework/dal"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/DataTable"
)

func GetManualListByUserId(userId int64) ([]*DataTable.UserManual, error) {
	sqlstr := dal.MarshalGetSql(DataTable.NewUserManual(), "user_id")
	rows, err := Control.GameDB.Query(sqlstr, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	manualList := make([]*DataTable.UserManual, 0)
	for rows.Next() {
		manual := DataTable.NewUserManual()
		err := rows.Scan(
			&manual.Id,
			&manual.UserId,
			&manual.ManualId,
			&manual.Level,
			&manual.UpdateTime,
			&manual.ActiveTime,
		)
		if err != nil {
			return nil, err
		}
		manualList = append(manualList, manual)
	}
	return manualList, nil
}
