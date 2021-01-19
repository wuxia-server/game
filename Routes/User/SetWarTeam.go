package User

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/Rule"
)

type SetWarTeam struct {
	Network.WebSocketRoute

	TeamId int // 队伍ID
}

func (e *SetWarTeam) Parse() {
	e.TeamId = utils.NewStringAny(e.Params["team_id"]).ToIntV()
}

func (e *SetWarTeam) Handle(agent *Network.WebSocketAgent) uint32 {
	person := Manage.GetPersonByAgent(agent)
	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	if person.User.WarTeamId == e.TeamId {
		return Code.User_SetWarTeam_AlreadyWar
	}

	if person.GetTeam(e.TeamId) == nil {
		return Code.User_SetWarTeam_TeamNotExists
	}

	person.User.WarTeamId = e.TeamId
	person.User.Save()

	e.Mod(Rule.RULE_USER, person.User.ToJsonMap())

	return messages.RC_Success
}
