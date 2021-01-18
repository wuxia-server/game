package Data

import (
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/DataTable"
)

// 队伍列表转为JsonMap输出格式
func (e *Person) __TeamToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range e.TeamList {
		result[utils.NewStringInt(k).ToString()] = v.ToJsonMap()
	}
	return result
}

func (e *Person) GetTeamById(teamId int) *DataTable.UserTeam {
	for _, val := range e.TeamList {
		if val.TeamId == teamId {
			return val
		}
	}
	return nil
}

func (e *Person) HaveTeam(teamId int) bool {
	return e.GetTeamById(teamId) != nil
}