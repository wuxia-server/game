package Data

import (
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/DataTable"
)

// 物品列表转为JsonMap输出格式
func (e *Person) __ItemToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range e.ItemList {
		result[utils.NewStringInt(k).ToString()] = v.Num
	}
	return result
}

func (e *Person) GetItemById(itemId int) *DataTable.UserItem {
	for _, item := range e.ItemList {
		if item.ItemId == itemId {
			return item
		}
	}
	return nil
}
