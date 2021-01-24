package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/utils"
	"time"
)

type Sign struct {
	Id     int `ST:"PK"`        // 唯一ID
	Mark   int `ST:"mark"`      // 标志 (1、2标识不断的循环使用月签到数据)
	Day    int `ST:"day"`       // 天数
	DropId int `ST:"sign_drop"` // 签到掉落ID
}

var (
	_SignList []*Sign
)

func init() {
	filePath := "./JSON/wx_sign.json"
	rows, err := Table.LoadTable(filePath, &Sign{})
	if err != nil {
		panic(err)
	}

	arr := make([]*Sign, 0)
	for _, row := range rows {
		arr = append(arr, row.(*Sign))
	}
	_SignList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetSign(day int) (result *Sign) {
	for _, row := range _SignList {
		if row.Day == day {
			newrow := utils.ReflectNew(row)
			result = newrow.(*Sign)
			break
		}
	}
	return
}

func GetSignNowDay() (result *Sign) {
	return GetSign(time.Now().Day())
}
