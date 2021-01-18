package main

import (
	"github.com/team-zf/framework"
	"github.com/team-zf/framework/Network"
	"github.com/team-zf/framework/config"
	"github.com/team-zf/framework/modules"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/Routes"
	_ "github.com/wuxia-server/game/StaticTable"
	"time"
)

func main() {
	Control.App = framework.CreateApp(
		modules.AppSetDebug(true),
		modules.AppSetParse(true),
		modules.AppSetTableDir("./JSON"),
		modules.AppSetPStatusTime(5*time.Second),
	)
	Control.App.OnConfigurationLoaded(func(app modules.IApp, conf *config.AppConfig) {
		// 载入数据库模块(账户服)
		if item := conf.Settings["gate_db"]; item != nil {
			settings := item.(map[string]interface{})
			Control.GateDB = modules.NewDataBaseModule(
				modules.DataBaseSetDsn(settings["dsn"].(string)),
			)
			app.AddModule(Control.GateDB)
		}
		// 载入数据库模块(逻辑服)
		if item := conf.Settings["game_db"]; item != nil {
			settings := item.(map[string]interface{})
			Control.GameDB = modules.NewDataBaseModule(
				modules.DataBaseSetDsn(settings["dsn"].(string)),
			)
			app.AddModule(Control.GameDB)
		}

		// 载入WS服务模块
		app.AddModule(Network.NewWebSocketModule(
			Network.WebSocketSetName("逻辑服"),
			Network.WebSocketSetAddr(":20301"),
			Network.WebSocketSetRoute(Routes.Route),
		))
	})
	Control.App.Init()
	Control.App.Run()
}
