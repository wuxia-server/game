package Code

import "github.com/wuxia-server/game/Cmd"

/**
 * 账户注册
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Client_Connect*100 + iota

	Client_Connect_TokenIncorrect   // Token错误
	Client_Connect_AccountDisable   // 该账户已被禁用 (也许是被封号了)
	Client_Connect_AlreadyConnected // 已经连接上了
)

/**
 * 创建角色
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.User_CreateRole*100 + iota

	User_CreateRole_GenerateUserIdFail  // UserId生成失败  (UserId生成次数达上限)
	User_CreateRole_NicknameHasBeenUsed // 昵称已被使用
	User_CreateRole_UserHasCreated      // 用户已经创建
)

/**
 * 设置出战队伍
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.User_SetWarTeam*100 + iota

	User_SetWarTeam_AlreadyWar    // 已经出战
	User_SetWarTeam_TeamNotExists // 该队伍不存在
)

/**
 * 设置队伍的英雄位置
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Team_SetPosition*100 + iota

	Team_SetPosition_TeamNotExists   // 没有这个队伍
	Team_SetPosition_PositionInvalid // 无效的位置
	Team_SetPosition_HeroNotExists   // 没有这个英雄
	Team_SetPosition_HeroAlreadyWar  // 该英雄已经出战了
)

/**
 * 佩戴玄气
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Hero_WearSpirit*100 + iota

	Hero_WearSpirit_SlotIncalid // 无效的槽位
)

/**
 * 饲养经验
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Hero_FeedExp*100 + iota

	Hero_FeedExp_ItemNotExists        // 没有这个物品
	Hero_FeedExp_HeroNotExists        // 没有这个英雄
	Hero_FeedExp_QuantityInsufficient // 数量不足
	Hero_FeedExp_HeroExpSpiledOver    // 该英雄经验已满
)
