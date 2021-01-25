package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
)

type Role struct {
	RoleId int `ST:"PK"`   // 角色ID
	ItemId int `ST:"gift"` // 赠送物品ID
}

var (
	_RoleList []*Role
)

func init() {
	filePath := "./JSON/wx_role.json"
	rows, err := Table.LoadTable(filePath, &Role{})
	if err != nil {
		panic(err)
	}

	arr := make([]*Role, 0)
	for _, row := range rows {
		arr = append(arr, row.(*Role))
	}
	_RoleList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetRole(roleId int) (result *Role) {
	for _, row := range _RoleList {
		if row.RoleId == roleId {
			result = row
			break
		}
	}
	return
}
