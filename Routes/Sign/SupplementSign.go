package Sign

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Const"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/StaticTable"
)

type SupplementSign struct {
	Network.WebSocketRoute

	Day int // 补签日期(日)
}

func (e *SupplementSign) Parse() {
	e.Day = utils.NewStringAny(e.Params["day"]).ToIntV()
}

func (e *SupplementSign) Handle(agent *Network.WebSocketAgent) uint32 {
	person := Manage.GetPersonByAgent(agent)
	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	stSign := StaticTable.GetSign(e.Day)
	// 没有这个签到日
	if stSign == nil {
		return Code.Sign_SupplementSign_UnableSign
	}

	dtSign := person.GetSign(stSign.Day)
	// 已经签到了
	if dtSign != nil && dtSign.Status > 0 {
		return Code.Sign_SupplementSign_AlreadySign
	}

	ddm, err := person.Sign(stSign.Day, Const.SignMethod_SupplementSign)
	// 已经签到了
	if err != nil {
		return Code.Sign_SupplementSign_AlreadySign
	}
	e.Join(ddm)

	// 处理掉落
	{
		gainItems, ddm := person.Drop2(stSign.DropId)
		e.Data("gain_items", gainItems)
		e.Join(ddm)
	}

	return messages.RC_Success
}
