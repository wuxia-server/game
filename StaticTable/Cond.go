package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type Cond struct {
	CondId     int         `ST:"PK"`         // 条件ID
	CondType   int         `ST:"cond_type"`  // 条件类型
	CondLogic  int         `ST:"cond_logic"` // 条件逻辑符
	CondParams *Table.List `ST:"cond_para"`  // 条件参数 (根据类型需要的参数可能有1个或者多个)
}

var (
	_CondList []*Cond
)

func init() {
	filePath := "./JSON/wx_cond.json"
	rows, err := Table.LoadTable(filePath, &Cond{})
	if err != nil {
		panic(err)
	}

	arr := make([]*Cond, 0)
	for _, row := range rows {
		arr = append(arr, row.(*Cond))
	}
	_CondList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetCond(condId int) (result *Cond) {
	for _, row := range _CondList {
		if row.CondId == condId {
			result = row
			break
		}
	}
	return
}
