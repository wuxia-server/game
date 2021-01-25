package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type Shop struct {
	ShopId            int         `ST:"PK"`           // 商店ID
	ShopType          int         `ST:"shop_type"`    // 商店类型 (1、兑换商店 2、活动商店)
	LimitCondIds      *Table.List `ST:"limit_cond"`   // 限制条件ID列表
	GoodsBank         int         `ST:"goods_bank"`   // 商品库
	RefreshGap        int         `ST:"refresh_gap"`  // 自动刷新的间隔时间 (秒)
	SupportRefresh    bool        `ST:"is_refresh"`   // 是否支持手动刷新
	RefreshItemCost   *Table.Map  `ST:"refresh_cost"` // 手动刷新物品消耗 (物品ID, 数量)
	RefreshTicketCost *Table.Map  `ST:"cost_ticket"`  // 手动刷新券消耗 (券ID, 数量)
	ValidTime         *Table.List `ST:"open_limit"`   // 有效期 (开始时间, 结束时间)
}

var (
	_ShopList []*Shop
)

func init() {
	filePath := "./JSON/wx_shop.json"
	rows, err := Table.LoadTable(filePath, &Shop{})
	if err != nil {
		panic(err)
	}

	arr := make([]*Shop, 0)
	for _, row := range rows {
		arr = append(arr, row.(*Shop))
	}
	_ShopList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetShop(shopId int) (result *Shop) {
	for _, row := range _ShopList {
		if row.ShopId == shopId {
			result = row
			break
		}
	}
	return
}
