package User

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Manage"
)

type SetWarTeam struct {
	Network.WebSocketRoute

	TeamId int // 队伍ID
}

func (e *SetWarTeam) Parse() {
	e.TeamId = utils.NewStringAny(e.Params["team_id"]).ToIntV()
}

func (e *SetWarTeam) Handle(agent *Network.WebSocketAgent) uint32 {
	client := Manage.GetPersonByAgent(agent)

	if client.User.WarTeamId == e.TeamId {
		return Code.User_SetWarTeam_AlreadyWar
	}

	if client.GetTeamById(e.TeamId) == nil {
		return Code.User_SetWarTeam_TeamNotExists
	}

	client.User.WarTeamId = e.TeamId
	client.User.Save()

	//client.User.ToJsonMap()

	return messages.RC_Success
}
