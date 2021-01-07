package Data

import (
	"github.com/team-zf/framework/dal"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/Table"
)

func GetUserByAccountId(accountId int64) (user *Table.User) {
	user = Table.NewUser()
	sqlstr := dal.MarshalGetSql(user, "account_id")
	row := Control.GameDB.QueryRow(sqlstr, accountId)
	if row.Scan(
		&user.UserId,
		&user.AccountId,
		&user.AvatarId,
		&user.Nickname,
		&user.Level,
		&user.Exp,
		&user.VipLevel,
		&user.VipExp,
		&user.Status,
		&user.WarTeamId,
		&user.Vigor,
		&user.VigorRecoverTime,
		&user.Vitality,
		&user.VitalityRecoverTime,
		&user.CreateTime,
		&user.OnlineTime,
		&user.OfflineTime,
	) != nil {
		user = nil
	}
	return
}
