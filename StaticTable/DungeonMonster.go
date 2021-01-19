package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/utils"
)

type DungeonMonster struct {
	MonsterId     int         `ST:"PK"`                   // 怪物ID
	MonsterLevel  int         `ST:"monster_lv"`           // 怪物等级
	MonsterTrain  *Table.List `ST:"monster_train"`        // 怪物养成
	HeroId        int         `ST:"model_hero_id"`        // 模板英雄ID
	TalismanId    int         `ST:"model_talisman_id"`    // 模板法宝ID
	TalismanTrain *Table.List `ST:"model_talisman_train"` // 模板法宝养成
	MountId       int         `ST:"model_mount_id"`       // 模板坐骑ID
	MountTrain    *Table.List `ST:"model_mount_train"`    // 模板坐骑养成
}

var (
	_DungeonMonsterList []*DungeonMonster
)

func init() {
	filePath := "./JSON/wx_dungeon_monster.json"
	rows, err := Table.LoadTable(filePath, &DungeonMonster{})
	if err != nil {
		panic(err)
	}

	arr := make([]*DungeonMonster, 0)
	for _, row := range rows {
		arr = append(arr, row.(*DungeonMonster))
	}
	_DungeonMonsterList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetDungeonMonster(monsterId int) (result *DungeonMonster) {
	for _, row := range _DungeonMonsterList {
		if row.MonsterId == monsterId {
			newrow := utils.ReflectNew(row)
			result = newrow.(*DungeonMonster)
			break
		}
	}
	return
}
