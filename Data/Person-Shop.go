package Data

import (
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/DataTable"
)

func (e *Person) __ShopToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range e.ShopList {
		result[utils.NewStringInt(k).ToString()] = v.ToJsonMap()
	}
	return result
}

func (e *Person) GetShop(shopId int) (result *DataTable.UserShop) {
	for _, shop := range e.ShopList {
		if shop.ShopId == shopId {
			result = shop
			break
		}
	}
	return
}
