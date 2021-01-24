package Dal

import (
	"github.com/team-zf/framework/dal"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/DataTable"
)

func GetShopListByUserId(userId int64) ([]*DataTable.UserShop, error) {
	sqlstr := dal.MarshalGetSql(DataTable.NewUserShop(), "user_id")
	rows, err := Control.GameDB.Query(sqlstr, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	shopList := make([]*DataTable.UserShop, 0)
	for rows.Next() {
		shop := DataTable.NewUserShop()
		str := ""
		err := rows.Scan(
			&shop.Id,
			&shop.UserId,
			&shop.ShopId,
			&str,
			&shop.UpdateTime,
		)
		if err != nil {
			return nil, err
		}
		shop.DetailList = DataTable.NewShopDetailListS(str)
		shopList = append(shopList, shop)
	}
	return shopList, nil
}
