package DataTable

import (
	"github.com/team-zf/framework/dal"
	"time"
)

type UserGoods struct {
	dal.BaseTable

	Id         int64           `db:"id,pk"`        // 主键 (用户ID+商店ID)
	UserId     int64           `db:"user_id,!mod"` // 用户ID
	ShopId     int             `db:"shop_id,!mod"` // 商店ID
	DetailList *ShopDetailList `db:"detail_list"`  // 细节列表
	UpdateTime time.Time       `db:"update_time"`  // 更新时间

	// Id      int64 `db:"id,pk"`        // 主键 (用户ID+商品)
	// UserId  int64 `db:"user_id,!mod"` // 用户ID
	// ShopId  int   `db:"shop_id"`      // 商店ID
	// GoodsId int   `db:"goods_id"`     // 商品ID
	// Sales   int   `db:"sales"`        // 销量
}
