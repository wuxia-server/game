package Manage

import (
	"github.com/wuxia-server/game/Client"
	"sync"
)

var (
	mutex   sync.RWMutex
	clients map[int64]*Client.ClientModel
)

func init() {
	clients = make(map[int64]*Client.ClientModel)
}

func AddClient(client *Client.ClientModel) {
	if client == nil {
		return
	}

	if client.Account == nil || client.Account.Id == 0 {
		return
	}

	mutex.Lock()
	clients[client.Account.Id] = client
	mutex.Unlock()
}
