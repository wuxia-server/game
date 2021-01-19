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

func (e *Person) GetHero(heroId int) (result *DataTable.UserHero) {
	for _, hero := range e.HeroList {
		if hero.HeroId == heroId {
			result = hero
			break
		}
	}
	return
}

func (e *Person) HaveHero(heroId int) bool {
	return e.GetHero(heroId) != nil
}
