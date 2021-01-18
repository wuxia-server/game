package Dungeon

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
)

type Attack struct {
	Network.WebSocketRoute

	StoryId int // 关卡ID
}

func (e *Attack) Parse() {
	e.StoryId = utils.NewStringAny(e.Params["story_id"]).ToIntV()
}

func (e *Attack) Handle(agent *Network.WebSocketAgent) uint32 {
	return messages.RC_Success
}
