package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/utils"
)

type ManualEducate struct {
	Level      int         `ST:"PK"`         // 图鉴等级
	Stage      int         `ST:"stage"`      // 图鉴阶段
	StageLevel int         `ST:"stage_lv"`   // 阶段等级
	Property   *Table.List `ST:"prop"`       // 养成属性
	CostSoul   int         `ST:"cost_soul"`  // 消耗魂魄
	ExtraCost  *Table.Map  `ST:"extra_cost"` // 额外消耗
}

var (
	_ManualEducateList []*ManualEducate
	_ManualLevelLimit  int
)

func init() {
	filePath := "./JSON/wx_manual_educate.json"
	rows, err := Table.LoadTable(filePath, &ManualEducate{})
	if err != nil {
		panic(err)
	}

	arr := make([]*ManualEducate, 0)
	for _, row := range rows {
		arr = append(arr, row.(*ManualEducate))
	}
	_ManualEducateList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetManualEducate(level int) (result *ManualEducate) {
	for _, row := range _ManualEducateList {
		if row.Level == level {
			newrow := utils.ReflectNew(row)
			result = newrow.(*ManualEducate)
			break
		}
	}
	return
}

func GetManualLevelLimit() int {
	if _ManualLevelLimit > 0 {
		return _ManualLevelLimit
	}
	levelLimit := 0
	for _, row := range _ManualEducateList {
		if row.Level > levelLimit {
			levelLimit = row.Level
		}
	}
	_ManualLevelLimit = levelLimit
	return levelLimit
}
