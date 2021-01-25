package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type Manual struct {
	ManualId   int `ST:"PK"`   // 图鉴ID
	ManualBank int `ST:"bank"` // 图鉴库
}

var (
	_ManualList []*Manual
)

func init() {
	filePath := "./JSON/wx_manual.json"
	rows, err := Table.LoadTable(filePath, &Manual{})
	if err != nil {
		panic(err)
	}

	arr := make([]*Manual, 0)
	for _, row := range rows {
		arr = append(arr, row.(*Manual))
	}
	_ManualList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetManual(manualId int) (result *Manual) {
	for _, row := range _ManualList {
		if row.ManualId == manualId {
			result = row
			break
		}
	}
	return
}
