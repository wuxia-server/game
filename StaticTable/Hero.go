package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type Hero struct {
	HeroId         int         `ST:"PK"`            // 英雄ID
	Race           int         `ST:"hero_race"`     // 种族(1、道 2、佛 3、妖 4、巫 5、人 6、主)
	Elements       int         `ST:"hero_elements"` // 五行(1、金 2、木 3、水 4、火 5、土)
	Quality        int         `ST:"hero_quality"`  // 品质
	Star           int         `ST:"hero_star"`     // 星级
	Design         int         `ST:"hero_design"`   // 定位(1、攻击型 2、防御型 3、辅助性)
	LinkHeroId     int         `ST:"link_hero"`     // 关联英雄ID
	InitPropery    *Table.List `ST:"init_prop"`     // 初始属性
	ProperyRate    *Table.List `ST:"prop_rate"`     // 属性万分比
	NatkSkillId    int         `ST:"normal_atk"`    // 普攻技能ID
	MainSkillId    int         `ST:"aid_atk"`       // 主动技能ID
	PassiveSkillId int         `ST:"main_atk"`      // 被动技能ID
	AidSkillId     int         `ST:"awake_skill"`   // 援助技能ID
}

var (
	_HeroList []*Hero
)

func init() {
	filePath := "./JSON/wx_hero.json"
	rows, err := Table.LoadTable(filePath, &Hero{})
	if err != nil {
		panic(err)
	}

	arr := make([]*Hero, 0)
	for _, row := range rows {
		arr = append(arr, row.(*Hero))
	}
	_HeroList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetHero(heroId int) (result *Hero) {
	for _, row := range _HeroList {
		if row.HeroId == heroId {
			result = row
			break
		}
	}
	return
}
