package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type ManualBank struct {
	Id             int         `ST:"PK"`                 // 唯一ID
	ManualBank     int         `ST:"bank"`               // 图鉴库
	ActivePropList *Table.List `ST:"bank_active_effect"` // 激活属性列表
	BankGroup      int         `ST:"bank_group"`         // 图鉴库组
}

var (
	_ManualBankList []*ManualBank
)

func init() {
	filePath := "./JSON/wx_manual_bank.json"
	rows, err := Table.LoadTable(filePath, &ManualBank{})
	if err != nil {
		panic(err)
	}

	arr := make([]*ManualBank, 0)
	for _, row := range rows {
		arr = append(arr, row.(*ManualBank))
	}
	_ManualBankList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetManualBank(id int) (result *ManualBank) {
	for _, row := range _ManualBankList {
		if row.Id == id {
			result = row
			break
		}
	}
	return
}
