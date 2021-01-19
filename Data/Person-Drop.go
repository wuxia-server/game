package Data

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/StaticTable"
)

func (e *Person) Drop(dropId int) ([]*DropItem, *Network.WebSocketDDM) {
	// 掉落细节列表
	dropDetails := make(map[int]*StaticTable.DropDetail)
	for _, drop := range StaticTable.GetDropList(dropId) {
		if utils.Percent(drop.Prob) {
			np := 0
			prob := utils.PercentV()
			for _, detail := range StaticTable.GetDropDetailList(drop.DropDetailId) {
				if detail.DropProb+np >= prob {
					dropDetails[detail.Id] = detail
					break
				}
				np += detail.DropProb
			}
		}
	}

	ddmSum := new(Network.WebSocketDDM)
	dropItemList := make([]*DropItem, 0)

	for _, detail := range dropDetails {
		arr, _ := detail.DropNum.ToIntArray()
		num := utils.Range(arr[0], arr[1])

		ddmSum.Join(e.AddItemV(detail.DropItemId, num))

		dropItem := new(DropItem)
		dropItem.ItemId = detail.DropItemId
		dropItem.Num = num
		dropItemList = append(dropItemList, dropItem)
	}

	return dropItemList, ddmSum
}

func (e *Person) Drop2(dropId int) ([][]int, *Network.WebSocketDDM) {
	dropItems, ddm := e.Drop(dropId)
	gainItems := make([][]int, 0)
	for _, dropItem := range dropItems {
		gainItems = append(gainItems, dropItem.ToArray())
	}
	return gainItems, ddm
}
