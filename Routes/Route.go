package Routes

import (
	"github.com/team-zf/framework/Network"
	"github.com/wuxia-server/game/Cmd"
	"github.com/wuxia-server/game/Routes/Cave"
	"github.com/wuxia-server/game/Routes/Client"
	"github.com/wuxia-server/game/Routes/Dungeon"
	"github.com/wuxia-server/game/Routes/Hero"
	"github.com/wuxia-server/game/Routes/Lottery"
	"github.com/wuxia-server/game/Routes/Manual"
	"github.com/wuxia-server/game/Routes/Shop"
	"github.com/wuxia-server/game/Routes/Sign"
	"github.com/wuxia-server/game/Routes/Task"
	"github.com/wuxia-server/game/Routes/Team"
	"github.com/wuxia-server/game/Routes/User"
)

var (
	Route = Network.NewWebSocketRouteHandle()
)

func init() {
	Route.SetRoute(Cmd.Client_Connect, &Client.Connect{})
	Route.SetRoute(Cmd.Client_Home, &Client.Home{})
	Route.SetRoute(Cmd.User_CreateRole, &User.CreateRole{})
	Route.SetRoute(Cmd.User_SetWarTeam, &User.SetWarTeam{})
	Route.SetRoute(Cmd.Team_SetPosition, &Team.SetPosition{})
	Route.SetRoute(Cmd.Hero_WearSpirit, &Hero.WearSpirit{})
	Route.SetRoute(Cmd.Hero_FeedExp, &Hero.FeedExp{})
	Route.SetRoute(Cmd.Manual_Upgrade, &Manual.Upgrade{})
	Route.SetRoute(Cmd.Shop_Refresh, &Shop.Refresh{})
	Route.SetRoute(Cmd.Shop_Buy, &Shop.Buy{})
	Route.SetRoute(Cmd.Sign_Sign, &Sign.Sign{})
	Route.SetRoute(Cmd.Sign_SupplementSign, &Sign.SupplementSign{})
	Route.SetRoute(Cmd.Task_Reward, &Task.Reward{})
	Route.SetRoute(Cmd.Lottery_Card, &Lottery.Card{})
	Route.SetRoute(Cmd.Cave_Harvest, &Cave.Harvest{})
	Route.SetRoute(Cmd.Cave_Upgrade, &Cave.Upgrade{})
	Route.SetRoute(Cmd.Dungeon_Attack, &Dungeon.Attack{})
}
