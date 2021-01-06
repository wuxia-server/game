package Code

import "github.com/wuxia-server/game/WebSocketRoute/Cmd"

/**
 * 账户注册
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Client_Connect*100 + iota

	Client_Connect_TokenIncorrect // Token错误
	Client_Connect_AccountDisable // 该账户已被禁用 (也许是被封号了)
)
