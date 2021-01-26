package Team

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/Rule"
)

type SetPosition struct {
	Network.WebSocketRoute

	TeamId   int // 队伍ID
	HeroId   int // 英雄ID
	Position int // 位置
}

func (e *SetPosition) Parse() {
	e.TeamId = utils.NewStringAny(e.Params["team_id"]).ToIntV()
	e.HeroId = utils.NewStringAny(e.Params["hero_id"]).ToIntV()
	e.Position = utils.NewStringAny(e.Params["pos"]).ToIntV()
}

func (e *SetPosition) Handle(agent *Network.WebSocketAgent) uint32 {
	person := Manage.GetPersonByAgent(agent)
	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	team := person.GetTeam(e.TeamId)
	// 没有这个队伍
	if team == nil {
		return Code.Team_SetPosition_TeamNotExists
	}

	// 无效的位置
	if e.Position < 1 || e.Position > 5 {
		return Code.Team_SetPosition_PositionInvalid
	}

	// 没有这个英雄
	if !person.HaveHero(e.HeroId) {
		return Code.Team_SetPosition_HeroNotExists
	}

	// 该英雄已经出战了
	if team.Position1 == e.HeroId || team.Position2 == e.HeroId || team.Position3 == e.HeroId || team.Position4 == e.HeroId || team.Position5 == e.HeroId {
		return Code.Team_SetPosition_HeroAlreadyWar
	}

	switch e.Position {
	case 1:
		team.Position1 = e.HeroId
	case 2:
		team.Position2 = e.HeroId
	case 3:
		team.Position3 = e.HeroId
	case 4:
		team.Position4 = e.HeroId
	case 5:
		team.Position5 = e.HeroId
	}
	team.Save()

	e.Mod(Rule.RULE_TEAM, team.ToJsonMap())
	return messages.RC_Success
}
