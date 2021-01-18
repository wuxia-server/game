package Dal

import (
	"github.com/team-zf/framework/dal"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/DataTable"
	"math/rand"
	"time"
)

func GetUserByAccountId(accountId int64) (user *DataTable.User) {
	user = DataTable.NewUser()
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

func GetUserByNickname(nickname string) (user *DataTable.User) {
	user = DataTable.NewUser()
	sqlstr := dal.MarshalGetSql(user, "name")
	row := Control.GameDB.QueryRow(sqlstr, nickname)
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

func GenerateUserId() (userId int64, ok bool) {
	retryCount := 0
	rand.Seed(time.Now().Unix())
	for {
		retryCount++
		// 重试次数达到100次以后直接宣布生成失败
		if retryCount > 100 {
			return 0, false
		}
		userId = rand.Int63n(89999) + 10000 // 随机生成一个七位数ID
		sqlstr := `select name from user where id = ?;`
		row := Control.GameDB.QueryRow(sqlstr, userId)
		username := ""
		if row.Scan(&username) != nil {
			break
		}
	}
	return userId, true
}
