package Data

import (
	"errors"
	"fmt"
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/DataTable"
	"github.com/wuxia-server/game/Rule"
)

// 英雄列表转为JsonMap输出格式
func (e *Person) __HeroToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for _, hero := range e.HeroList {
		result[utils.NewStringInt(hero.HeroId).ToString()] = hero.ToJsonMap()
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

func (e *Person) AddHero(heroId int) (*Network.WebSocketDDM, error) {
	hero := e.GetHero(heroId)
	if hero != nil {
		return nil, errors.New(fmt.Sprintf("用户(%d)已经拥有HeroId(%d), 无法重复获得.", e.UserId(), heroId))
	}

	hero = DataTable.NewUserHero()
	hero.Id = e.JoinToUserId(heroId)
	hero.UserId = e.UserId()
	hero.HeroId = heroId
	hero.Level = 1
	hero.Exp = 0
	hero.Evolution = 0
	hero.Latent1 = 0
	hero.Latent2 = 0
	hero.Latent3 = 0
	hero.Latent4 = 0
	hero.TalismanSlot = 0
	hero.MountSlot = 0
	hero.SpiritSlot1 = 0
	hero.SpiritSlot2 = 0
	hero.SpiritSlot3 = 0
	hero.SpiritSlot4 = 0
	hero.FightPower = 0
	hero.Save()

	if e.HeroList == nil {
		e.HeroList = []*DataTable.UserHero{hero}
	} else {
		e.HeroList = append(e.HeroList, hero)
	}

	ddm := new(Network.WebSocketDDM)
	ddm.Mod(Rule.RULE_HERO, hero.ToJsonMap())
	return ddm, nil
}
