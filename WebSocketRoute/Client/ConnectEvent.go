package Client

import (
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/model"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Client"
	"github.com/wuxia-server/game/Data"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/WebSocketRoute/Code"
)

type ConnectEvent struct {
	messages.HttpMessage

	Token string // 登录口令
}

func (e *ConnectEvent) Parse() {
	e.Token = utils.NewStringAny(e.Params["token"]).ToString()
}

func (e *ConnectEvent) WebSocketDirectCall(wsmd *model.WebSocketModel, resp *messages.WebSocketResponse) {
	account := Data.GetAccountByToken(e.Token)

	// Token错误
	if account == nil {
		logger.Debug("Token错误")
		resp.Code = Code.Client_Connect_TokenIncorrect
		return
	}

	// 该账户已被禁用
	if account.Status == 1 {
		logger.Debug("该账户已被禁用")
		resp.Code = Code.Client_Connect_AccountDisable
	}

	client := Client.NewClientModel(wsmd)
	client.Account = account
	Manage.AddClient(client)

	user := Data.GetUserByAccountId(account.Id)
	resp.Data["is_new"] = user == nil

	if user != nil {
		client.User = user
		client.UpdateVigorRecover()
		client.UpdateVitalityRecover()
	}

	logger.Debug("连接成功")

	resp.Code = messages.RC_Success
}

func M_Connect() *ConnectEvent {
	return &ConnectEvent{}
}
