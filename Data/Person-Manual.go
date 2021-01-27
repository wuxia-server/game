package Data

import (
	"errors"
	"fmt"
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/DataTable"
	"github.com/wuxia-server/game/Rule"
	"time"
)

// 图鉴列表转为JsonMap输出格式
func (e *Person) __ManualToJson() map[string]interface{} {
	result := make(map[string]interface{})
	for _, manual := range e.ManualList {
		result[utils.NewStringInt(manual.ManualId).ToString()] = manual.Level
	}
	return result
}

func (e *Person) GetManual(manualId int) (result *DataTable.UserManual) {
	for _, manual := range e.ManualList {
		if manual.ManualId == manualId {
			result = manual
			break
		}
	}
	return
}

func (e *Person) AddManual(manualId int) (*Network.WebSocketDDM, error) {
	manual := e.GetManual(manualId)
	if manual != nil {
		return nil, errors.New(fmt.Sprintf("用户(%d)已经拥有ManualId(%d), 无法重复获得.", e.UserId(), manualId))
	}

	manual = DataTable.NewUserManual()
	manual.Id = e.JoinToUserId(manualId)
	manual.UserId = e.UserId()
	manual.ManualId = manualId
	manual.Level = 1
	manual.ActiveTime = time.Now()
	manual.UpdateTime = time.Now()
	manual.Save()

	if e.ManualList == nil {
		e.ManualList = []*DataTable.UserManual{manual}
	} else {
		e.ManualList = append(e.ManualList, manual)
	}

	ddm := new(Network.WebSocketDDM)
	ddm.Mod(Rule.RULE_MANUAL, manual.ToJsonMap())
	return ddm, nil
}

func (e *Person) UpgradeManual(manualId int) (*Network.WebSocketDDM, error) {
	manual := e.GetManual(manualId)
	if manual == nil {
		return nil, errors.New(fmt.Sprintf("用户(%d)没有ManualId(%d), 无法升级.", e.UserId(), manualId))
	}

	manual.Level = manual.Level + 1
	manual.UpdateTime = time.Now()

	ddm := new(Network.WebSocketDDM)
	ddm.Mod(Rule.RULE_MANUAL, manual.ToJsonMap())
	return ddm, nil
}
