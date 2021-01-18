package DataTable

import (
	"github.com/team-zf/framework/dal"
)

type UserHero struct {
	dal.BaseTable

	Id           int64 `db:"id,pk"`         // 主键 (用户ID+英雄ID)
	UserId       int   `db:"user_id,!mod"`  // 用户ID
	HeroId       int   `db:"hero_id,!mod"`  // 英雄ID
	Level        int   `db:"level"`         // 等级
	Exp          int   `db:"exp"`           // 经验
	Evolution    int   `db:"evolution"`     // 渡劫阶段
	Latent1      int   `db:"latent_1"`      // 潜力点 1
	Latent2      int   `db:"latent_2"`      // 潜力点 2
	Latent3      int   `db:"latent_3"`      // 潜力点 3
	Latent4      int   `db:"latent_4"`      // 潜力点 4
	TalismanSlot int   `db:"talisman_slot"` // 法宝槽位
	MountSlot    int   `db:"mount_slot"`    // 坐骑槽位
	SpiritSlot1  int   `db:"spirit_slot_1"` // 玄气槽位 (血量)
	SpiritSlot2  int   `db:"spirit_slot_2"` // 玄气槽位 (攻击)
	SpiritSlot3  int   `db:"spirit_slot_3"` // 玄气槽位 (法术)
	SpiritSlot4  int   `db:"spirit_slot_4"` // 玄气槽位 (速度)
	FightPower   int   `db:"fight_power"`   // 总战斗力
}

func (e *UserHero) GetTableName() (name string) {
	return "user_hero"
}

func (e *UserHero) ToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["hero_id"] = e.HeroId
	result["level"] = e.Level
	result["exp"] = e.Exp
	result["evolution"] = e.Evolution
	result["latent_1"] = e.Latent1
	result["latent_2"] = e.Latent2
	result["latent_3"] = e.Latent3
	result["latent_4"] = e.Latent4
	result["talisman_slot"] = e.TalismanSlot
	result["mount_slot"] = e.MountSlot
	result["spirit_slot_1"] = e.SpiritSlot1
	result["spirit_slot_2"] = e.SpiritSlot2
	result["spirit_slot_3"] = e.SpiritSlot3
	result["spirit_slot_4"] = e.SpiritSlot4
	result["fight_power"] = e.FightPower
	return result
}

func NewUserHero() *UserHero {
	result := new(UserHero)
	result.BaseTable.Init(result)
	return result
}
