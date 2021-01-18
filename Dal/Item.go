package Dal

import (
	"github.com/team-zf/framework/dal"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/DataTable"
)

func GetItemListByUserId(userId int64) ([]*DataTable.UserItem, error) {
	sqlstr := dal.MarshalGetSql(DataTable.NewUserItem(), "user_id")
	rows, err := Control.GameDB.Query(sqlstr, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	itemList := make([]*DataTable.UserItem, 0)
	for rows.Next() {
		item := DataTable.NewUserItem()
		err := rows.Scan(
			&item.Id,
			&item.UserId,
			&item.ItemId,
			&item.Num,
			&item.UpdateTime,
		)
		if err != nil {
			return nil, err
		}
		itemList = append(itemList, item)
	}
	return itemList, nil
}
