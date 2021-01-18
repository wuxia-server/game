package Dal

import (
	"github.com/team-zf/framework/dal"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/DataTable"
)

func GetHeroListByUserId(userId int64) ([]*DataTable.UserHero, error) {
	sqlstr := dal.MarshalGetSql(DataTable.NewUserHero(), "user_id")
	rows, err := Control.GameDB.Query(sqlstr, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	heroList := make([]*DataTable.UserHero, 0)
	for rows.Next() {
		hero := DataTable.NewUserHero()
		err := rows.Scan(
			&hero.Id,
			&hero.UserId,
			&hero.HeroId,
			&hero.Level,
			&hero.Exp,
			&hero.Evolution,
			&hero.Latent1,
			&hero.Latent2,
			&hero.Latent3,
			&hero.Latent4,
			&hero.TalismanSlot,
			&hero.MountSlot,
			&hero.SpiritSlot1,
			&hero.SpiritSlot2,
			&hero.SpiritSlot3,
			&hero.SpiritSlot4,
			&hero.FightPower,
		)
		if err != nil {
			return nil, err
		}
		heroList = append(heroList, hero)
	}
	return heroList, nil
}
