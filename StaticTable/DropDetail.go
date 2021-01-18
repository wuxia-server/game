package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/utils"
)

type DropDetail struct {
	Id           int         `ST:"PK"`             // 唯一ID
	DropDetailId int         `ST:"drop_detail_id"` // 掉落细节ID
	DropType     int         `ST:"drop_type"`      // 掉落类型 (1.数值 2.普通)
	DropItemId   int         `ST:"drop_items_id"`  // 掉落物品ID
	DropNum      *Table.List `ST:"drop_num"`       // 掉落数量
	DropProb     int         `ST:"drop_prob"`      // 掉落概率 (万分比)
}

var (
	_DropDetailList []*DropDetail
)

func init() {
	filePath := "./JSON/wx_drop_detail.json"
	rows, err := Table.LoadTable(filePath, &DropDetail{})
	if err != nil {
		panic(err)
	}

	arr := make([]*DropDetail, 0)
	for _, row := range rows {
		arr = append(arr, row.(*DropDetail))
	}
	_DropDetailList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetDropDetail(id int) (result *DropDetail) {
	for _, row := range _DropDetailList {
		if row.Id == id {
			newrow := utils.ReflectNew(row)
			result = newrow.(*DropDetail)
			break
		}
	}
	return
}

func GetDropDetailList(dropDetailId int) (result []*DropDetail) {
	result = make([]*DropDetail, 0)
	for _, row := range _DropDetailList {
		if row.DropDetailId == dropDetailId {
			newrow := utils.ReflectNew(row)
			result = append(result, newrow.(*DropDetail))
		}
	}
	return
}
