package Dal

import (
	"github.com/team-zf/framework/dal"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/DataTable"
)

func GetAccountByToken(token string) (account *DataTable.Account) {
	account = DataTable.NewAccount()
	sqlstr := dal.MarshalGetSql(account, "token")
	row := Control.GateDB.QueryRow(sqlstr, token)
	if row.Scan(
		&account.Id,
		&account.UserName,
		&account.PassWord,
		&account.Email,
		&account.Phone,
		&account.Token,
		&account.Status,
		&account.LatelyServer,
		&account.CreateTime,
	) != nil {
		account = nil
	}
	return
}
