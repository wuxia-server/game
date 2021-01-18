package Client

import (
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Dal"
	"github.com/wuxia-server/game/Data"
	"github.com/wuxia-server/game/Manage"
)

type Connect struct {
	Network.WebSocketRoute

	Token string
}

func (e *Connect) Parse() {
	e.Token = utils.NewStringAny(e.Params["token"]).ToString()
}

func (e *Connect) Handle(agent *Network.WebSocketAgent) uint32 {
	// 已经连接上了
	if Manage.GetPersonByAgent(agent) != nil {
		return Code.Client_Connect_AlreadyConnected
	}

	account := Dal.GetAccountByToken(e.Token)

	// Token错误
	if account == nil {
		return Code.Client_Connect_TokenIncorrect
	}

	// 该账户已被禁用
	if account.Status == 1 {
		return Code.Client_Connect_AccountDisable
	}

	person := new(Data.Person)
	person.Agent = agent
	person.Account = account
	Manage.AddPerson(person)

	user := Dal.GetUserByAccountId(account.Id)
	if user != nil {
		person.User = user
		person.UpdateVigorRecover()
		person.UpdateVitalityRecover()
	}
	e.Data("is_new", user == nil)

	return messages.RC_Success
}
