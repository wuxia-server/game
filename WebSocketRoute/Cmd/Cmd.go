package Cmd

/**
 * 客户端模块
 */
const (
	Client_Connect uint32 = 1001 // 连接
	Client_Home    uint32 = 1002 // 主页
)

/**
 * 用户模块
 */
const (
	User_CreateRole uint32 = 1101 // 创建角色
	User_WarTeam    uint32 = 1102 // 设置出战队伍
)

/**
 * 队伍模块
 */
const (
	Team_Position uint32 = 1201 // 设置队伍的英雄位置
)

/**
 * 英雄模块
 */
const (
	Hero_WearSpirit uint32 = 1301 // 佩戴玄气
	Hero_FeedExp    uint32 = 1302 // 饲养经验
)

/**
 * 图鉴模块
 */
const (
	Manual_Upgrade uint32 = 1401 // 升级
)

/**
 * 商店模块
 */
const (
	Shop_Refresh uint32 = 1501 // 刷新
	Shop_Buy     uint32 = 1502 // 购买
)

/**
 * 签到模块
 */
const (
	Sign_Sign           uint32 = 1601 // 签到
	Sign_SupplementSign uint32 = 1602 // 补签
)

/**
 * 签到模块
 */
const (
	Task_Reward uint32 = 1701 // 领取任务奖励
)

/**
 * 抽奖模块
 */
const (
	Lottery_Card uint32 = 1801 // 抽卡
)

/**
 * 洞府模块
 */
const (
	Cave_Harvest uint32 = 1901 // 丰收
	Cave_Upgrade uint32 = 1902 // 升级
)

/**
 * 关卡模块
 */
const (
	Dungeon_Attack uint32 = 2001 // 攻打
)
