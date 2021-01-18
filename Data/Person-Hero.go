package Data

import (
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/DataTable"
)

// 英雄列表转为JsonMap输出格式
func (e *Person) __HeroToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range e.HeroList {
		result[utils.NewStringInt(k).ToString()] = v.ToJsonMap()
	}
	return result
}

func (e *Person) GetHeroById(heroId int) *DataTable.UserHero {
	for _, hero := range e.HeroList {
		if hero.HeroId == heroId {
			return hero
		}
	}
	return nil
}

func (e *Person) HaveHero(heroId int) bool {
	return e.GetHeroById(heroId) != nil
}
