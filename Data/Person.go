package Data

import (
	"fmt"
	"github.com/team-zf/framework/Network"
	"github.com/wuxia-server/game/Dal"
	"github.com/wuxia-server/game/DataTable"
	"strconv"
	"sync"
)

type Person struct {
	mutex sync.Mutex
	Agent *Network.WebSocketAgent

	Account     *DataTable.Account
	User        *DataTable.User
	HeroList    []*DataTable.UserHero    // 英雄列表
	ItemList    []*DataTable.UserItem    // 物品列表
	TeamList    []*DataTable.UserTeam    // 队伍列表
	CaveList    []*DataTable.UserCave    // 洞府列表
	SignList    []*DataTable.UserSign    // 签到列表
	ManualList  []*DataTable.UserManual  // 图鉴列表
	DungeonList []*DataTable.UserDungeon // 副本列表
}

func (e *Person) UserId() int64 {
	if e.User == nil {
		return 0
	}
	return e.User.UserId
}

func (e *Person) AccountId() int64 {
	if e.Account == nil {
		return 0
	}
	return e.Account.Id
}

func (e *Person) JoinToUserId(id int) int64 {
	str := fmt.Sprintf("%d%d", e.UserId(), id)
	v, _ := strconv.ParseInt(str, 10, 64)
	return v
}

func (e *Person) Load() {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	if len(e.HeroList) == 0 {
		heroList, err := Dal.GetHeroListByUserId(e.UserId())
		if err != nil {
			panic(err)
		}
		e.HeroList = heroList
	}
	if len(e.ItemList) == 0 {
		itemList, err := Dal.GetItemListByUserId(e.UserId())
		if err != nil {
			panic(err)
		}
		e.ItemList = itemList
	}
	if len(e.TeamList) == 0 {
		teamList, err := Dal.GetTeamListByUserId(e.UserId())
		if err != nil {
			panic(err)
		}
		e.TeamList = teamList
	}
	if len(e.CaveList) == 0 {
		caveList, err := Dal.GetCaveListByUserId(e.UserId())
		if err != nil {
			panic(err)
		}
		e.CaveList = caveList
	}
	if len(e.SignList) == 0 {
		signList, err := Dal.GetSignListByUserId(e.UserId())
		if err != nil {
			panic(err)
		}
		e.SignList = signList
	}
	if len(e.ManualList) == 0 {
		manualList, err := Dal.GetManualListByUserId(e.UserId())
		if err != nil {
			panic(err)
		}
		e.ManualList = manualList
	}
	if len(e.DungeonList) == 0 {
		dungeonList, err := Dal.GetDungeonListByUserId(e.UserId())
		if err != nil {
			panic(err)
		}
		e.DungeonList = dungeonList
	}
}

func (e *Person) ToJsonMap() map[string]interface{} {
	return map[string]interface{}{
		"user":    e.__UserToJsonMap(),
		"hero":    e.__HeroToJsonMap(),
		"team":    e.__TeamToJsonMap(),
		"item":    e.__ItemToJsonMap(),
		"cave":    e.__CaveToJsonMap(),
		"manual":  e.__ManualToJson(),
		"sign":    e.__SignToJsonMap(),
		"dungeon": e.__DungeonToJsonMap(),
	}
}
