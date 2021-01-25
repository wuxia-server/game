package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type Drop struct {
	Id           int `ST:"PK"`         // 唯一ID
	DropId       int `ST:"item_type"`  // 掉落ID
	DropDetailId int `ST:"use_limit"`  // 掉落细节ID
	Prob         int `ST:"sell_state"` // 概率值(万分比)
}

var (
	_DropList []*Drop
)

func init() {
	filePath := "./JSON/wx_drop.json"
	rows, err := Table.LoadTable(filePath, &Drop{})
	if err != nil {
		panic(err)
	}

	arr := make([]*Drop, 0)
	for _, row := range rows {
		arr = append(arr, row.(*Drop))
	}
	_DropList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetDrop(id int) (result *Drop) {
	for _, row := range _DropList {
		if row.Id == id {
			result = row
			break
		}
	}
	return
}

func GetDropList(dropId int) (result []*Drop) {
	result = make([]*Drop, 0)
	for _, row := range _DropList {
		if row.DropId == dropId {
			result = append(result, row)
		}
	}
	return
}
