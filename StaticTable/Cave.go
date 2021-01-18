package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/utils"
)

type Cave struct {
	CaveId       int `ST:"PK"`          // 洞府ID
	RewardItemId int `ST:"reward_item"` // 产出物品ID
	RewardTime   int `ST:"reward_time"` // 可收获的最小生产时间 (秒)
	UpgradeId    int `ST:"upgrade"`     // 等级ID
}

var (
	_CaveList []*Cave
)

func init() {
	filePath := "./JSON/wx_cave.json"
	rows, err := Table.LoadTable(filePath, &Cave{})
	if err != nil {
		panic(err)
	}

	arr := make([]*Cave, 0)
	for _, row := range rows {
		arr = append(arr, row.(*Cave))
	}
	_CaveList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetCave(caveId int) (result *Cave) {
	for _, row := range _CaveList {
		if row.CaveId == caveId {
			newrow := utils.ReflectNew(row)
			result = newrow.(*Cave)
			break
		}
	}
	return
}

func GetCaveList() (result []*Cave) {
	result = make([]*Cave, 0)
	for _, row := range _CaveList {
		newrow := utils.ReflectNew(row)
		result = append(result, newrow.(*Cave))
	}
	return
}
