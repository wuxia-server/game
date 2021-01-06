@echo off

cd ../

set app_path=bin\pc\wuxia_game.exe
set conf_path=bin\pc\server.json

if not exist %app_path% (
	echo Need to build first!
	pause
	exit
)

if not exist %conf_path% (
	echo Need to build first!
	pause
	exit
)

%app_path% -c %conf_path%
