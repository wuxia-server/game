package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type ManualBankDetail struct {
	Id         int         `ST:"PK"`          // 唯一ID
	BankGroup  int         `ST:"bank_group"`  // 图鉴库组
	ActiveType int         `ST:"active_type"` // 激活类型
	ActiveId   int         `ST:"active_id"`   // 激活ID
	PropList   *Table.List `ST:"prop"`        // 图鉴属性系数
	SoulId     int         `ST:"soul_id"`     // 魂魄ID
}

var (
	_ManualBankDetailList []*ManualBankDetail
)

func init() {
	filePath := "./JSON/wx_manual_bank_detail.json"
	rows, err := Table.LoadTable(filePath, &ManualBankDetail{})
	if err != nil {
		panic(err)
	}

	arr := make([]*ManualBankDetail, 0)
	for _, row := range rows {
		arr = append(arr, row.(*ManualBankDetail))
	}
	_ManualBankDetailList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetManualBankDetail(id int) (result *ManualBankDetail) {
	for _, row := range _ManualBankDetailList {
		if row.Id == id {
			result = row
			break
		}
	}
	return
}
