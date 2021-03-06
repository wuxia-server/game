@echo off

set dest_path=..\bin

if not exist %dest_path% (
	md %dest_path%
)

set dest_path=%dest_path%\pc

if not exist %dest_path% (
	md %dest_path%
)

set app_path=%dest_path%\wuxia_game.exe
set _app_path=..\main.go
set conf_path=%dest_path%\server.json
set _conf_path=..\server.json

if exist %app_path% (
	del %app_path%
)

if exist %conf_path% (
	del %conf_path%
)

set GOARCH=amd64
set GOOS=windows

echo Building Project ...
echo Windows

go build -o %app_path% %_app_path%
copy %_conf_path% %conf_path%

echo Build Complete !
echo .
echo .
echo .

pause
