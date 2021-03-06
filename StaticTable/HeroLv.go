package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type HeroLv struct {
	Level    int         `ST:"PK"`       // 等级
	NeedExp  int         `ST:"need_exp"` // 需要的经验
	PropList *Table.List `ST:"prop"`     // 对应的属性
}

var (
	_HeroLvList     []*HeroLv
	_HeroLevelLimit int
)

func init() {
	filePath := "./JSON/wx_hero_lv.json"
	rows, err := Table.LoadTable(filePath, &HeroLv{})
	if err != nil {
		panic(err)
	}

	arr := make([]*HeroLv, 0)
	for _, row := range rows {
		arr = append(arr, row.(*HeroLv))
	}
	_HeroLvList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetHeroLv(level int) (result *HeroLv) {
	for _, row := range _HeroLvList {
		if row.Level == level {
			result = row
			break
		}
	}
	return
}

func GetHeroLvByExp(exp int) (result *HeroLv) {
	for _, row := range _HeroLvList {
		if row.NeedExp <= exp {
			if result == nil || result.Level < row.Level {
				result = row
			}
		}
	}
	return
}

func GetHeroLevelLimit() int {
	if _HeroLevelLimit > 0 {
		return _HeroLevelLimit
	}
	levelLimit := 0
	for _, row := range _HeroLvList {
		if row.Level > levelLimit {
			levelLimit = row.Level
		}
	}
	_HeroLevelLimit = levelLimit
	return levelLimit
}
