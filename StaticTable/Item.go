package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/utils"
)

type Item struct {
	ItemId       int         `ST:"PK"`             // 物品ID
	ItemType     int         `ST:"item_type"`      // 物品类型 (1.货币 2.英雄魂魄 3.英雄经验道具 4.普通材料)
	UseCondIds   *Table.List `ST:"use_limit"`      // 使用条件ID列表
	SupportSell  bool        `ST:"sell_state"`     // 是否支持卖出
	SellGetItems *Table.List `ST:"sell_get_items"` // 卖出获得物品列表
	GetWays      *Table.List `ST:"get_ways"`       // 获得途径列表
	ReelateValue int         `ST:"related_func"`   // 关联功能
}

var (
	_ItemList []*Item
)

func init() {
	filePath := "./JSON/wx_item.json"
	rows, err := Table.LoadTable(filePath, &Item{})
	if err != nil {
		panic(err)
	}

	arr := make([]*Item, 0)
	for _, row := range rows {
		arr = append(arr, row.(*Item))
	}
	_ItemList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetItem(itemId int) (result *Item) {
	for _, row := range _ItemList {
		if row.ItemId == itemId {
			newrow := utils.ReflectNew(row)
			result = newrow.(*Item)
			break
		}
	}
	return
}
