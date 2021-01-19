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

// 物品列表转为JsonMap输出格式
func (e *Person) __ItemToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range e.ItemList {
		result[utils.NewStringInt(k).ToString()] = v.Num
	}
	return result
}

func (e *Person) GetItem(itemId int) (result *DataTable.UserItem) {
	for _, item := range e.ItemList {
		if item.ItemId == itemId {
			result = item
			break
		}
	}
	return
}

// 增加物品数量
func (e *Person) AddItem(itemId int, num int) (*Network.WebSocketDDM, error) {
	if num <= 0 {
		return nil, errors.New(fmt.Sprintf("传入的数量有误(%d).", num))
	}

	item := e.GetItem(itemId)
	if item == nil {
		item = new(DataTable.UserItem)
		item.Id = e.JoinToUserId(itemId)
		item.UserId = e.UserId()
		item.ItemId = itemId
		item.Num = num
		item.UpdateTime = time.Now()
		e.ItemList[item.ItemId] = item
	} else {
		item.Num += num
	}
	item.Save()

	ddm := new(Network.WebSocketDDM)
	ddm.Mod(Rule.RULE_ITEM, item.ToJsonMap())
	return ddm, nil
}

// 增加物品数量(异常直接抛panic)
func (e *Person) AddItemV(itemId int, num int) *Network.WebSocketDDM {
	ddm, err := e.AddItem(itemId, num)
	if err != nil {
		panic(err)
	}
	return ddm
}

// 减少物品数量
func (e *Person) SubItem(itemId int, num int) (*Network.WebSocketDDM, error) {
	if num <= 0 {
		return nil, errors.New(fmt.Sprintf("传入的数量有误(%d).", num))
	}

	item := e.GetItem(itemId)
	if item == nil {
		return nil, errors.New(fmt.Sprintf("没有物品(%d).", itemId))
	}

	if item.Num < num {
		return nil, errors.New(fmt.Sprintf("需要消耗数量(%d), 物品(%d)数量不足.", num, itemId))
	}

	item.Num -= num
	item.Save()

	ddm := new(Network.WebSocketDDM)
	ddm.Mod(Rule.RULE_ITEM, item.ToJsonMap())
	return ddm, nil
}

// 减少物品数量(异常直接抛panic)
func (e *Person) SubItemV(itemId int, num int) *Network.WebSocketDDM {
	ddm, err := e.SubItem(itemId, num)
	if err != nil {
		panic(err)
	}
	return ddm
}

// 批量减少物品数量
func (e *Person) SubItems(costs map[int]int) (*Network.WebSocketDDM, error) {
	for itemId, num := range costs {
		item := e.GetItem(itemId)
		if item == nil {
			return nil, errors.New(fmt.Sprintf("没有物品(%d).", itemId))
		}
		if item.Num < num {
			return nil, errors.New(fmt.Sprintf("需要消耗数量(%d), 物品(%d)数量不足.", num, itemId))
		}
	}

	ddm := new(Network.WebSocketDDM)
	for itemId, num := range costs {
		item := e.GetItem(itemId)
		item.Num -= num
		item.Save()
		ddm.Mod(Rule.RULE_ITEM, item.ToJsonMap())
	}
	return ddm, nil
}

// 批量减少物品数量(异常直接抛panic)
func (e *Person) SubItemsV(costs map[int]int) *Network.WebSocketDDM {
	ddm, err := e.SubItems(costs)
	if err != nil {
		panic(err)
	}
	return ddm
}
