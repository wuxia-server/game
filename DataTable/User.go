package DataTable

import (
	"github.com/team-zf/framework/dal"
	"github.com/team-zf/framework/logger"
	"github.com/wuxia-server/game/Control"
	"time"
)

type User struct {
	dal.BaseTable

	UserId              int64     `db:"user_id,pk"`            // 主键
	AccountId           int64     `db:"account_id,!mod"`       // 账户关联ID
	AvatarId            int       `db:"avatar_id"`             // 头像ID
	Nickname            string    `db:"name"`                  // 用户昵称
	Level               int       `db:"level"`                 // 等级
	Exp                 int       `db:"exp"`                   // 经验
	VipLevel            int       `db:"vip_level"`             // VIP等级
	VipExp              int       `db:"vip_exp"`               // VIP经验
	Status              int       `db:"status"`                // 状态(0:正常 1:被封)
	WarTeamId           int       `db:"war_team_id"`           // 出战队伍ID
	Vigor               int       `db:"vigor"`                 // 体力值
	VigorRecoverTime    time.Time `db:"vigor_recover_time"`    // 体力值最后更新时间(用于实现定时增长)
	Vitality            int       `db:"vitality"`              // 活力值
	VitalityRecoverTime time.Time `db:"vitality_recover_time"` // 活力值最后更新时间(用于实现定时增长)
	CreateTime          time.Time `db:"create_time"`           // 创建时间
	OnlineTime          time.Time `db:"online_time"`           // 最后一次上线时间
	OfflineTime         time.Time `db:"offline_time"`          // 最后一次下线时间
}

func (e *User) ToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["user_id"] = e.UserId
	result["avatar_id"] = e.AvatarId
	result["nick_name"] = e.Nickname
	result["level"] = e.Level
	result["exp"] = e.Exp
	result["vip_level"] = e.VipLevel
	result["vip_exp"] = e.VipExp
	result["war_team_id"] = e.WarTeamId
	result["vigor"] = e.Vigor
	result["vigor_recover_time"] = e.VigorRecoverTime.Unix()
	result["vitality"] = e.Vitality
	result["vitality_recover_time"] = e.VitalityRecoverTime.Unix()
	return result
}

func (e *User) Save() {
	msg := &dal.DalMessage{
		UserId: e.UserId,
		Table:  e,
	}
	msg.RunFunc = func(db dal.IConnDB) error {
		_, err := db.Exec(dal.MarshalModSql(e),
			e.UserId, e.AccountId, e.AvatarId,
			e.Nickname, e.Level, e.Exp, e.VipLevel, e.VipExp, e.Status, e.WarTeamId,
			e.Vigor, e.VigorRecoverTime, e.Vitality, e.VitalityRecoverTime,
			e.CreateTime, e.OnlineTime, e.OfflineTime,
		)
		if err != nil {
			logger.Error("数据表[%s]保存失败, 错误原因: %+v", e.GetTableName(), err)
		}
		return err
	}
	Control.GameDB.AddMsg(msg)
}

func NewUser() *User {
	result := new(User)
	result.BaseTable.Init(result)
	return result
}
