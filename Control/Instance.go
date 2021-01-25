package Control

import (
	"github.com/team-zf/framework/DB"
	"github.com/team-zf/framework/modules"
)

var (
	App    modules.IApp
	GateDB DB.IDataBaseModule
	GameDB DB.IDataBaseModule
)
