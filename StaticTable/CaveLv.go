package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type CaveLv struct {
	Id      int         `ST:"PK"`      // 等级ID
	Level   int         `ST:"level"`   // 等级
	Speed   int         `ST:"speed"`   // 产量/时
	Maximum int         `ST:"maximum"` // 最大容量
	NextId  int         `ST:"next_lv"` // 下一等级ID
	Cost    *Table.List `ST:"cost"`    // 升级消耗 (物品ID, 数量)
}

var (
	_CaveLvList []*CaveLv
)

func init() {
	filePath := "./JSON/wx_cave_upgrade.json"
	rows, err := Table.LoadTable(filePath, &CaveLv{})
	if err != nil {
		panic(err)
	}

	arr := make([]*CaveLv, 0)
	for _, row := range rows {
		arr = append(arr, row.(*CaveLv))
	}
	_CaveLvList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetCaveLv(id int) (result *CaveLv) {
	for _, row := range _CaveLvList {
		if row.Id == id {
			result = row
			break
		}
	}
	return
}

func GetCaveLvList() (result []*CaveLv) {
	result = make([]*CaveLv, 0)
	for _, row := range _CaveLvList {
		result = append(result, row)
	}
	return
}

func GetCaveLvNext(id int) (result *CaveLv) {
	row := GetCaveLv(id)
	if row != nil && row.NextId > 0 {
		result = GetCaveLv(row.NextId)
	}
	return
}
