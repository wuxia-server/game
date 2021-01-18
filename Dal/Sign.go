package Dal

import (
	"github.com/team-zf/framework/dal"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/DataTable"
)

func GetSignListByUserId(userId int64) ([]*DataTable.UserSign, error) {
	sqlstr := dal.MarshalGetSql(DataTable.NewUserSign(), "user_id")
	rows, err := Control.GameDB.Query(sqlstr, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	signList := make([]*DataTable.UserSign, 0)
	for rows.Next() {
		sign := DataTable.NewUserSign()
		err := rows.Scan(
			&sign.Id,
			&sign.UserId,
			&sign.Day,
			&sign.Status,
			&sign.SignTime,
		)
		if err != nil {
			return nil, err
		}
		signList = append(signList, sign)
	}
	return signList, nil
}
