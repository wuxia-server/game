package User

import (
	"fmt"
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/messages"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Code"
	"github.com/wuxia-server/game/Const"
	"github.com/wuxia-server/game/Dal"
	"github.com/wuxia-server/game/DataTable"
	"github.com/wuxia-server/game/Manage"
	"github.com/wuxia-server/game/StaticTable"
	"time"
)

type CreateRole struct {
	Network.WebSocketRoute

	RoleId   int    // 主角ID
	Nickname string // 指定昵称
}

func (e *CreateRole) Parse() {
	e.RoleId = utils.NewStringAny(e.Params["role_id"]).ToIntV()
	e.Nickname = utils.NewStringAny(e.Params["nickname"]).ToString()

	if StaticTable.GetRole(e.RoleId) == nil {
		panic(fmt.Sprintf("未找到RoleId(%d)", e.RoleId))
	}
}

func (e *CreateRole) Handle(agent *Network.WebSocketAgent) uint32 {
	person := Manage.GetPersonByAgent(agent)

	// 没有权限; 未连接
	if person == nil {
		return messages.RC_NoPermission
	}

	if Dal.GetUserByNickname(e.Nickname) != nil {
		return Code.User_CreateRole_NicknameHasBeenUsed
	}

	if Dal.GetUserByAccountId(person.AccountId()) != nil {
		return Code.User_CreateRole_UserHasCreated
	}

	userId, ok := Dal.GenerateUserId()
	if !ok {
		return Code.User_CreateRole_GenerateUserIdFail
	}

	// 用户信息
	{
		user := DataTable.NewUser()
		user.UserId = userId
		user.AccountId = person.AccountId()
		user.Nickname = e.Nickname
		user.WarTeamId = 1
		user.Level = 1
		user.Vigor = Const.InitialVigor
		user.VigorRecoverTime = time.Now()
		user.Vitality = Const.InitialVitality
		user.VitalityRecoverTime = time.Now()
		user.CreateTime = time.Now()
		user.OnlineTime = time.Now()
		user.OfflineTime = time.Now()
		user.Save()
		person.User = user
	}
	// 用户队伍
	{
		teamList := make([]*DataTable.UserTeam, 0)
		row := StaticTable.GetRoleLv(person.User.Level)
		for id := 1; id <= row.MaxTeamNum; id++ {
			team := DataTable.NewUserTeam()
			team.Id = person.JoinToUserId(id)
			team.UserId = person.UserId()
			team.TeamId = id
			team.FightPower = 0
			if id == 1 {
				team.DefendDongfu = 1
				team.DefendLundao = 1
			} else {
				team.DefendDongfu = 0
				team.DefendLundao = 0
			}
			team.Save()
			teamList = append(teamList, team)
		}
		person.TeamList = teamList
	}
	// 用户物品
	{
		itemList := make([]*DataTable.UserItem, 0)

		items := map[int]int{
			1011:   100000, // 银币
			1021:   50000,  // 元魂丹
			1031:   68888,  // 灵玉
			400101: 500,    // 飞升丹
		}
		// 英雄魂魄
		items[StaticTable.GetRole(e.RoleId).ItemId] = 1

		for key, value := range items {
			item := DataTable.NewUserItem()
			item.Id = person.JoinToUserId(key)
			item.UserId = person.UserId()
			item.ItemId = key
			item.Num = value
			item.UpdateTime = time.Now()
			item.Save()
			itemList = append(itemList, item)
		}
		person.ItemList = itemList
	}
	// 用户洞府
	{
		caveList := make([]*DataTable.UserCave, 0)
		for _, row := range StaticTable.GetCaveList() {
			cave := DataTable.NewUserCave()
			cave.Id = person.JoinToUserId(row.CaveId)
			cave.UserId = person.UserId()
			cave.CaveId = row.CaveId
			cave.UpgradeId = row.UpgradeId
			cave.UpdateTime = time.Now()
			cave.LastReceiveTime = time.Now()
			cave.Save()
			caveList = append(caveList, cave)
		}
		person.CaveList = caveList
	}
	// 用户副本
	{
		dungeonList := make([]*DataTable.UserDungeon, 0)
		chapter := StaticTable.GetDungeonChapterFirst()
		dungeon := DataTable.NewUserDungeon()
		dungeon.Id = person.JoinToUserId(chapter.ChapterId)
		dungeon.UserId = person.UserId()
		dungeon.StoryId = chapter.NormalStoryId
		dungeon.Save()
		dungeonList = append(dungeonList, dungeon)
		person.DungeonList = dungeonList
	}

	return messages.RC_Success
}
