package Client

import (
	"github.com/team-zf/framework/model"
	"github.com/wuxia-server/game/Table"
)

type ClientModel struct {
	wsmd    *model.WebSocketModel
	Account *Table.Account
	User    *Table.User
}

func (e *ClientModel) UserId() int64 {
	return e.User.UserId
}

func (e *ClientModel) AccountId() int64 {
	return e.Account.Id
}

func NewClientModel(wsmd *model.WebSocketModel) *ClientModel {
	return &ClientModel{wsmd: wsmd}
}
