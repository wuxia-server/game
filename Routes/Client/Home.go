package Client

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/wuxia-server/game/Manage"
	"time"
)

type Home struct {
	Network.WebSocketRoute
}

func (e *Home) Parse() {
}

func (e *Home) Handle(agent *Network.WebSocketAgent) uint32 {
	person := Manage.GetPersonByAgent(agent)

	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	person.Load()

	for key, val := range person.ToJsonMap() {
		e.Data(key, val)
	}
	e.Data("sys_time", time.Now().Unix())

	return messages.RC_Success
}
