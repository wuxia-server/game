package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type RoleLv struct {
	Level      int `ST:"PK"`            // 等级
	NeedExp    int `ST:"need_exp"`      // 需要的经验
	MaxTeamNum int `ST:"max_group_num"` // 最大队伍数量
}

var (
	_RoleLvList     []*RoleLv
	_RoleLevelLimit int
)

func init() {
	filePath := "./JSON/wx_role_lv.json"
	rows, err := Table.LoadTable(filePath, &RoleLv{})
	if err != nil {
		panic(err)
	}

	arr := make([]*RoleLv, 0)
	for _, row := range rows {
		arr = append(arr, row.(*RoleLv))
	}
	_RoleLvList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetRoleLv(level int) (result *RoleLv) {
	for _, row := range _RoleLvList {
		if row.Level == level {
			result = row
			break
		}
	}
	return
}

func GetRoleLevelLimit() int {
	if _RoleLevelLimit > 0 {
		return _RoleLevelLimit
	}
	levelLimit := 0
	for _, row := range _RoleLvList {
		if row.Level > levelLimit {
			levelLimit = row.Level
		}
	}
	_RoleLevelLimit = levelLimit
	return levelLimit
}
