package Data

import (
	"github.com/team-zf/framework/Network"
	"github.com/wuxia-server/game/StaticTable"
)

func (e *Person) CondVerify() *Network.WebSocketDDM {
	ddm := new(Network.WebSocketDDM)

	ddm.Join(e.__CondVerify_Task())
	ddm.Join(e.__CondVerify_Shop())

	return ddm
}

func (e *Person) __CondVerify_Task() *Network.WebSocketDDM {
	// 任务细节ID链表 (索引 0.当前任务细节ID  1.父级任务细节ID)
	dids := make([][]int, 0)
	for _, task := range StaticTable.GetTaskList() {
		ids, _ := task.DetailIds.ToIntArray()
		for _, id := range ids {
			dids = append(dids, []int{id})
			for {
				detail := StaticTable.GetTaskDetail(id)
				if detail.NextDetailId == -1 {
					break
				}
				id = detail.NextDetailId
				dids = append(dids, []int{detail.NextDetailId, detail.DetailId})
			}
		}
	}

	ddmSum := new(Network.WebSocketDDM)
	for _, v := range dids {
		// 子级
		if len(v) >= 1 && v[0] > 0 {
			ut := e.GetTask(v[0])
			if ut != nil {
				continue
			}
		}

		// 父级
		if len(v) >= 2 && v[1] > 0 {
			ut := e.GetTask(v[1])
			// 如果有父级, 要求先领取完
			if ut == nil || ut.Status != 2 {
				continue
			}
		}

		detail := StaticTable.GetTaskDetail(v[0])
		if e.__Verify(detail.TaskCondId) {
			ddm, _ := e.AddTask(detail.DetailId)
			ddmSum.Join(ddm)
		}
	}
	return ddmSum
}

func (e *Person) __CondVerify_Shop() *Network.WebSocketDDM {
	ddmSum := new(Network.WebSocketDDM)

	newShopList := make(map[int]int)
	for _, shop := range StaticTable.GetShopList() {
		if _, ok := newShopList[shop.ShopType]; ok {
			continue
		}
		ids, _ := shop.LimitCondIds.ToIntArray()
		if len(ids) == 0 || e.__VerifyIds(ids) {
			newShopList[shop.ShopType] = shop.ShopId
		}
	}

	oldShopList := make(map[int]int)
	for _, v := range e.ShopList {
		shop := StaticTable.GetShop(v.ShopId)
		oldShopList[shop.ShopType] = v.ShopId
	}

	for shopType, shopId := range oldShopList {
		if _, ok := newShopList[shopType]; ok {
			continue
		}
		ddm, err := e.DelShop(shopId)
		if err != nil {
			panic(err)
		} else {
			ddmSum.Join(ddm)
		}
	}

	for shopType, shopId := range newShopList {
		v, ok := oldShopList[shopType]
		if ok && v == shopId {
			continue
		}
		ddm, err := e.AddShop(shopId)
		if err != nil {
			panic(err)
		} else {
			ddmSum.Join(ddm)
		}
	}
	return ddmSum
}

// 效验条件ID (支持多个ID, 条件为并且)
func (e *Person) __Verify(args ...int) bool {
	ids := make([]int, 0)
	for _, id := range args {
		ids = append(ids, id)
	}
	return e.__VerifyIds(ids)
}

// 效验条件ID (仅支持ID数组格式)
func (e *Person) __VerifyIds(ids []int) bool {
	for _, id := range ids {
		r := false
		cond := StaticTable.GetCond(id)
		if cond == nil {
			return false
		}
		switch cond.CondType {
		// 玩家等级
		case 1:
			params, _ := cond.CondParams.ToIntArray()
			r = e.__CondType1(cond.CondLogic, params)
		// 玩家通关指定章节
		case 2:
			params, _ := cond.CondParams.ToIntArray()
			r = e.__CondType2(cond.CondLogic, params)
		// 玩家等级
		case 3:
			params, _ := cond.CondParams.ToIntArray()
			r = e.__CondType3(cond.CondLogic, params)
		}
		if !r {
			return false
		}
	}
	return true
}

// 玩家等级
func (e *Person) __CondType1(ls int, params []int) bool {
	return e.__LogicSymbol(ls, []int{e.Level()}, params)
}

// 玩家通关指定章节
func (e *Person) __CondType2(ls int, params []int) bool {
	// 通关章节ID列表 (仅普通关卡)
	perfect := make([]int, 0)
	for _, v := range StaticTable.GetDungeonChapterList() {
		// 是否通关
		result := true
		storyId := v.NormalStoryId
		for {
			StStory := StaticTable.GetDungeonStory(storyId)
			if StStory == nil {
				break
			}
			story := e.GetDungeon(storyId)
			if story == nil || story.Star < 1 {
				result = false
				break
			}
			storyId = StStory.NextId
		}

		if result {
			perfect = append(perfect, v.ChapterId)
		}
	}
	return e.__LogicSymbol(ls, perfect, params)
}

// 玩家获得指定英雄
func (e *Person) __CondType3(ls int, params []int) bool {
	arr := make([]int, 0)
	for _, v := range e.HeroList {
		arr = append(arr, v.HeroId)
	}
	return e.__LogicSymbol(ls, arr, params)
}

// 各种类型的逻辑运算符
func (e *Person) __LogicSymbol(ls int, curr []int, need []int) bool {
	switch ls {
	// 等于
	case 1:
		{
			for _, v := range curr {
				if v == need[0] {
					return true
				}
			}
		}
	// 大于
	case 2:
		{
			for _, v := range curr {
				if v > need[0] {
					return true
				}
			}
		}
	// 小于
	case 3:
		{
			for _, v := range curr {
				if v < need[0] {
					return true
				}
			}
		}
	// 不等于
	case 4:
		{
			for _, v := range curr {
				if v != need[0] {
					return true
				}
			}
		}
	// 大于等于
	case 5:
		{
			for _, v := range curr {
				if v >= need[0] {
					return true
				}
			}
		}
	// 小于等于
	case 6:
		{
			for _, v := range curr {
				if v <= need[0] {
					return true
				}
			}
		}
	// 在参数区间[min,max]
	case 7:
		{
			for _, v := range curr {
				if v >= need[0] && v <= need[1] {
					return true
				}
			}
		}
	}
	return false
}
