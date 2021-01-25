package Data

import (
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/DataTable"
)

// 队伍列表转为JsonMap输出格式
func (e *Person) __TeamToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for _, team := range e.TeamList {
		result[utils.NewStringInt(team.TeamId).ToString()] = team.ToJsonMap()
	}
	return result
}

func (e *Person) GetTeam(teamId int) (result *DataTable.UserTeam) {
	for _, team := range e.TeamList {
		if team.TeamId == teamId {
			result = team
			break
		}
	}
	return
}

func (e *Person) HaveTeam(teamId int) bool {
	return e.GetTeam(teamId) != nil
}
