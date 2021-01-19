package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/utils"
)

type LotteryDetail struct {
	Id          int         `ST:"PK"`           // 唯一ID
	LotteryBank int         `ST:"lottery_bank"` // 抽奖库
	DropId      int         `ST:"lottery_drop"` // 奖品掉落ID
	DropProb    int         `ST:"drop_weight"`  // 掉落权重
	DropMark    *Table.List `ST:"drop_mark"`    // 标记
}

var (
	_LotteryDetailList []*LotteryDetail
)

func init() {
	filePath := "./JSON/wx_lottery_detail.json"
	rows, err := Table.LoadTable(filePath, &LotteryDetail{})
	if err != nil {
		panic(err)
	}

	arr := make([]*LotteryDetail, 0)
	for _, row := range rows {
		arr = append(arr, row.(*LotteryDetail))
	}
	_LotteryDetailList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetLotteryDetail(id int) (result *LotteryDetail) {
	for _, row := range _LotteryDetailList {
		if row.Id == id {
			newrow := utils.ReflectNew(row)
			result = newrow.(*LotteryDetail)
			break
		}
	}
	return
}
