package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type Lottery struct {
	LotteryId    int         `ST:"PK"`           // 抽奖ID
	LotteryType  int         `ST:"lottery_type"` // 抽奖类型 (1、银币抽 2、灵玉抽)
	LotteryCount int         `ST:"lottery_time"` // 抽奖次数
	LotteryCost  *Table.Map  `ST:"lottery_cost"` // 抽奖花费
	LotteryBank  int         `ST:"lottery_bank"` // 抽奖库
	SpecailMark  int         `ST:"specail_mark"` // 特殊标记
	FakeDropList *Table.List `ST:"fake_drop"`    // 假掉落ID列表
}

var (
	_LotteryList []*Lottery
)

func init() {
	filePath := "./JSON/wx_lottery.json"
	rows, err := Table.LoadTable(filePath, &Lottery{})
	if err != nil {
		panic(err)
	}

	arr := make([]*Lottery, 0)
	for _, row := range rows {
		arr = append(arr, row.(*Lottery))
	}
	_LotteryList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetLottery(lotteryId int) (result *Lottery) {
	for _, row := range _LotteryList {
		if row.LotteryId == lotteryId {
			result = row
			break
		}
	}
	return
}
