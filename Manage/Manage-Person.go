package Manage

import (
	"errors"
	"github.com/team-zf/framework/Network"
	"github.com/wuxia-server/game/Data"
	"sync"
)

var (
	mutex sync.RWMutex
	/**
	 * key: 账户ID(AccoundId)
	 * val: Person
	 */
	_PersonList map[int64]*Data.Person
)

func init() {
	_PersonList = make(map[int64]*Data.Person)
}

func AddPerson(person *Data.Person) error {
	if person == nil || person.AccountId() == 0 {
		return errors.New("无法加入无效的Person.")
	}

	mutex.Lock()
	_PersonList[person.AccountId()] = person
	mutex.Unlock()
	return nil
}

func GetPersonByAgent(agent *Network.WebSocketAgent) *Data.Person {
	mutex.RLock()
	defer mutex.RUnlock()

	for _, person := range _PersonList {
		if person.Agent == agent {
			return person
		}
	}
	return nil
}
