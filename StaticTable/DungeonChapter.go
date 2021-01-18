package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Const"
)

type DungeonChapter struct {
	ChapterId     int `ST:"PK"`     // 章节ID
	NormalStoryId int `ST:"normal"` // 普通关卡ID
	EliteStoryId  int `ST:"elite"`  // 精英关卡ID
	NextChapterId int `ST:"next"`   // 下一个章节ID
}

var (
	_DungeonChapterList []*DungeonChapter
)

func init() {
	filePath := "./JSON/wx_dungeon_chapter.json"
	rows, err := Table.LoadTable(filePath, &DungeonChapter{})
	if err != nil {
		panic(err)
	}

	arr := make([]*DungeonChapter, 0)
	for _, row := range rows {
		arr = append(arr, row.(*DungeonChapter))
	}
	_DungeonChapterList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetDungeonChapter(chapterId int) (result *DungeonChapter) {
	for _, row := range _DungeonChapterList {
		if row.ChapterId == chapterId {
			newrow := utils.ReflectNew(row)
			result = newrow.(*DungeonChapter)
			break
		}
	}
	return
}

func GetDungeonChapterFirst() (result *DungeonChapter) {
	return GetDungeonChapter(Const.InitialChapterId)
}

func GetDungeonChapterList() (result []*DungeonChapter) {
	result = make([]*DungeonChapter, 0)
	for _, row := range _DungeonChapterList {
		newrow := utils.ReflectNew(row)
		result = append(result, newrow.(*DungeonChapter))
	}
	return
}
