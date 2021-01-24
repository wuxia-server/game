package Task

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/Rule"
	"github.com/wuxia-server/game/StaticTable"
)

type Reward struct {
	Network.WebSocketRoute

	DetailId int // 任务细节ID
}

func (e *Reward) Parse() {
	e.DetailId = utils.NewStringAny(e.Params["detail_id"]).ToIntV()
}

func (e *Reward) Handle(agent *Network.WebSocketAgent) uint32 {
	person := Manage.GetPersonByAgent(agent)
	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	stTaskDetail := StaticTable.GetTaskDetail(e.DetailId)
	// 没有这个任务细节
	if stTaskDetail == nil {
		return Code.Task_Reward_DetailNotExists
	}

	dtTask := person.GetTask(e.DetailId)
	// 未满足条件, 无法领取奖励
	if dtTask == nil {
		return Code.Task_Reward_UnableReward
	}

	// 已经领取过了
	if dtTask.Status == 2 {
		return Code.Task_Reward_AlreadyReward
	}

	// ### 处理领取
	{
		dtTask.Status = 2
		dtTask.Save()
		e.Mod(Rule.RULE_TASK, dtTask.ToJsonMap())
	}

	// ### 处理掉落
	{
		gainItems, ddm := person.Drop2(stTaskDetail.TaskDropId)
		e.Data("gain_items", gainItems)
		e.Join(ddm)
	}

	return messages.RC_Success
}
