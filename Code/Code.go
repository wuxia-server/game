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
	Hero_FeedExp_InvalidItem          // 无效的物品
	Hero_FeedExp_HeroNotExists        // 没有这个英雄
	Hero_FeedExp_QuantityInsufficient // 数量不足
	Hero_FeedExp_HeroExpFull          // 这个英雄的经验已经满了
)

/**
 * 丰收洞府
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Cave_Harvest*100 + iota

	Cave_Harvest_CaveNotExists     // 没有这个洞府
	Cave_Harvest_NotYetHarvestTime // 没到丰收时间
)

/**
 * 升级洞府
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Cave_Upgrade*100 + iota

	Cave_Upgrade_CaveNotExists    // 没有这个洞府
	Cave_Upgrade_AlreadyFullLevel // 无法升级, 已满级或达到了角色等级
	Cave_Upgrade_UnderCost        // 洞府升级成本不足
)

/**
 * 副本攻打
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Dungeon_Attack*100 + iota

	Dungeon_Attack_StoryNotExists        // 没有这个关卡
	Dungeon_Attack_NoRights              // 没有权限攻打
	Dungeon_Attack_AttackNumInsufficient // 挑战次数不足
	Dungeon_Attack_VigorInsufficient     // 体力不足
)

/**
 * 抽卡
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Lottery_Card*100 + iota

	Lottery_Card_LotteryNotExists // 没有这个抽奖项
	Lottery_Card_UnderCost        // 抽奖成本不足
)

/**
 * 图鉴升级
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Manual_Upgrade*100 + iota

	Manual_Upgrade_DetailNotExists // 没有这个图鉴细节
	Manual_Upgrade_LevelUpperLimit // 图鉴等级已达上限
	Manual_Upgrade_UnderCost       // 图鉴升级成本不足
)

/**
 * 商店刷新
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Shop_Refresh*100 + iota

	Shop_Refresh_ShopNotExists     // 沒有这个商店
	Shop_Refresh_ShopNotOpen       // 沒有开通这个商店
	Shop_Refresh_NotSupportRefresh // 该商店不支持手动刷新
	Shop_Refresh_UnderCost         // 商店刷新成本不足
)

/**
 * 商品购买
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Shop_Buy*100 + iota

	Shop_Buy_ShopNotExists   // 沒有这个商店
	Shop_Buy_DetailNotExists // 没有这个商品
	Shop_Buy_UnderCanbuy     // 商品购买次数不足
	Shop_Buy_UnderCost       // 商品购买成本不足
)

/**
 * 签到
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Sign_Sign*100 + iota

	Sign_Sign_UnableSign  // 没有这个签到日
	Sign_Sign_AlreadySign // 已经签到了
)

/**
 * 补签
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Sign_SupplementSign*100 + iota

	Sign_SupplementSign_UnableSign  // 没有这个签到日
	Sign_SupplementSign_AlreadySign // 已经签到了
)

/**
 * 领取任务奖励
 * 规则: (CMD*100) + iota
 * 例如: CMD=1001; CODE=100101
 */
const (
	_ = Cmd.Task_Reward*100 + iota

	Task_Reward_DetailNotExists // 没有这个任务细节
	Task_Reward_UnableReward    // 未满足条件, 无法领取奖励
	Task_Reward_AlreadyReward   // 已经领取过了
)
