package Data

import (
	"errors"
	"fmt"
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/DataTable"
	"github.com/wuxia-server/game/Rule"
	"time"
)

func (e *Person) __ShopToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for _, shop := range e.ShopList {
		result[utils.NewStringInt(shop.ShopId).ToString()] = shop.ToJsonMap()
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

func (e *Person) AddShop(shopId int) (*Network.WebSocketDDM, error) {
	shop := e.GetShop(shopId)
	if shop != nil {
		return nil, errors.New(fmt.Sprintf("用户(%d)已经拥有ShopId(%d), 无法重复获得.", e.UserId(), shopId))
	}

	shop = DataTable.NewUserShop()
	shop.Id = e.JoinToUserId(shopId)
	shop.UserId = e.UserId()
	shop.ShopId = shopId
	shop.UpdateTime = time.Now()
	shop.GenerateGoodsList()
	shop.Save()

	if e.ShopList == nil {
		e.ShopList = []*DataTable.UserShop{shop}
	} else {
		e.ShopList = append(e.ShopList, shop)
	}

	ddm := new(Network.WebSocketDDM)
	ddm.Mod(Rule.RULE_SHOP, shop.ToJsonMap())
	return ddm, nil
}

func (e *Person) DelShop(shopId int) (*Network.WebSocketDDM, error) {
	var idx int
	var shop *DataTable.UserShop
	for i, s := range e.ShopList {
		if s.ShopId == shopId {
			idx = i
			shop = s
			break
		}
	}

	if shop == nil {
		return nil, errors.New(fmt.Sprintf("用户(%d)没有ShopId(%d), 无法移除商店.", e.UserId(), shopId))
	}

	shop.Del()
	e.ShopList = append(e.ShopList[:idx], e.ShopList[idx+1:]...)

	ddm := new(Network.WebSocketDDM)
	ddm.Del(Rule.RULE_SHOP, shop.ToJsonMap())
	return ddm, nil
}
