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

// 签到列表转为JsonMap输出格式
func (e *Person) __SignToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for _, sign := range e.SignList {
		result[utils.NewStringInt(sign.Day).ToString()] = sign.ToJsonMap()
	}
	return result
}

func (e *Person) GetSign(day int) (result *DataTable.UserSign) {
	for _, sign := range e.SignList {
		if sign.Day == day {
			result = sign
			break
		}
	}
	return
}

/**
 * 签到
 * @param day   指定签到日
 * @param method 签到方式 (1=正常签到 2=补签)
 */
func (e *Person) Sign(day int, method int) (*Network.WebSocketDDM, error) {
	sign := e.GetSign(day)
	if sign != nil && sign.Status > 0 {
		return nil, errors.New(fmt.Sprintf("用户(%d)Day(%d)已经签过到了(%d), 无法重复签到.", e.UserId(), day, sign.Status))
	}

	sign = DataTable.NewUserSign()
	sign.Id = e.JoinToUserId(day)
	sign.UserId = e.UserId()
	sign.Day = day
	sign.Status = method
	sign.SignTime = time.Now()
	e.SignList = append(e.SignList, sign)

	sign.Save()

	ddm := new(Network.WebSocketDDM)
	ddm.Mod(Rule.RULE_SIGN, sign.ToJsonMap())
	return ddm, nil
}
