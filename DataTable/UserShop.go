package DataTable

import (
	"encoding/json"
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/dal"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/utils"
	"github.com/wuxia-server/game/Control"
	"github.com/wuxia-server/game/StaticTable"
	"strconv"
	"time"
)

type ShopDetailList struct {
	Items map[int]*ShopDetail
}

func NewShopDetailListS(str string) *ShopDetailList {
	data := make(map[int]map[string]interface{})
	json.Unmarshal([]byte(str), &data)

	result := new(ShopDetailList)
	result.Items = make(map[int]*ShopDetail)

	for k, v := range data {
		result.Items[k] = NewShopDetail(
			utils.NewStringAny(v["detail_id"]).ToIntV(),
			utils.NewStringAny(v["sales"]).ToIntV(),
			v["discount_params"],
		)
	}
	return result
}

func NewShopDetailList() *ShopDetailList {
	result := new(ShopDetailList)
	result.Items = make(map[int]*ShopDetail)
	return result
}

func (e *ShopDetailList) ToString() string {
	bytes, _ := json.Marshal(e.ToJsonMap())
	return string(bytes)
}

func (e *ShopDetailList) ToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range e.Items {
		result[strconv.Itoa(k)] = v.ToJsonMap()
	}
	return result
}

func (e *ShopDetailList) GetDetailById(detailId int) *ShopDetail {
	for _, v := range e.Items {
		if v.DetailId == detailId {
			return v
		}
	}
	return nil
}

type ShopDetail struct {
	DetailId       int         `db:"detail_id"`       // 细节ID
	Sales          int         `db:"sales"`           // 销量
	DiscountParams interface{} `db:"discount_params"` // 折扣参数
}

func (e *ShopDetail) ToJsonMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["detail_id"] = e.DetailId
	result["sales"] = e.Sales
	if e.DiscountParams != nil {
		switch e.DiscountParams.(type) {
		case *Table.Map:
			m := e.DiscountParams.(*Table.Map)
			result["discount_params"] = []int{
				m.KeyToInt(),
				m.ValueToInt(),
			}
		case int, float64:
			result["discount_params"] = utils.NewStringAny(e.DiscountParams).ToIntV()
		}
	}
	return result
}

func NewShopDetail(detailId int, sales int, discountParams interface{}) *ShopDetail {
	result := new(ShopDetail)
	result.DetailId = detailId
	result.Sales = sales
	result.DiscountParams = discountParams
	return result
}

type UserShop struct {
	dal.BaseTable

	Id         int64           `db:"id,pk"`        // 主键 (用户ID+商店ID)
	UserId     int64           `db:"user_id,!mod"` // 用户ID
	ShopId     int             `db:"shop_id,!mod"` // 商店ID
	DetailList *ShopDetailList `db:"detail_list"`  // 细节列表
	UpdateTime time.Time       `db:"update_time"`  // 更新时间
}

func (e *UserShop) GetTableName() (name string) {
	return "user_shop"
}

func (e *UserShop) ToJsonMap() map[string]interface{} {
	return map[string]interface{}{
		"shop_id":     e.ShopId,
		"detail_list": e.DetailList.ToJsonMap(),
		"update_time": e.UpdateTime.Unix(),
	}
}

func (e *UserShop) Save() {
	msg := &dal.DalMessage{
		UserId: e.UserId,
		Table:  e,
	}
	msg.RunFunc = func(db dal.IConnDB) error {
		_, err := db.Exec(dal.MarshalModSql(e),
			e.Id, e.UserId, e.ShopId,
			e.DetailList.ToString(), e.UpdateTime,
		)
		if err != nil {
			logger.Error("数据表[%s]保存失败, 错误原因: %+v", e.GetTableName(), err)
		}
		return err
	}
	Control.GameDB.AddMsg(msg)
}

func (e *UserShop) Del() {
	msg := &dal.DalMessage{
		UserId: e.UserId,
		Table:  e,
	}
	msg.RunFunc = func(db dal.IConnDB) error {
		_, err := db.Exec(dal.MarshalDelSql(e, "id"),
			e.Id,
		)
		if err != nil {
			logger.Error("数据表[%s]删除失败, 错误原因: %+v", e.GetTableName(), err)
		}
		return err
	}
	Control.GameDB.AddMsg(msg)
}

func (e *UserShop) GenerateGoodsList() {
	shop := StaticTable.GetShop(e.ShopId)
	details := StaticTable.GetShopDetailList(shop.GoodsBank)

	// 多组多项
	groups := make(map[int][]*StaticTable.ShopDetail)
	for _, v := range details {
		if groups[v.GoodsRank] == nil {
			groups[v.GoodsRank] = make([]*StaticTable.ShopDetail, 0)
		}
		groups[v.GoodsRank] = append(groups[v.GoodsRank], v)
	}

	// 多组单项
	group := make(map[int]*StaticTable.ShopDetail)
	for rank, arr := range groups {
		np := 0
		prob := utils.PercentV()
		for _, v := range arr {
			if v.GoodsRankProb+np >= prob {
				group[rank] = v
				break
			}
			np += v.GoodsRankProb
		}
	}

	// 更新商品列表
	e.DetailList = NewShopDetailList()
	for rank, v := range group {
		if v.DiscountParams.Empty() {
			e.DetailList.Items[rank] = NewShopDetail(v.Id, 0, nil)
		} else {
			e.DetailList.Items[rank] = NewShopDetail(v.Id, 0, v.DiscountParams.Rnd())
		}
	}
}

func NewUserShop() *UserShop {
	result := new(UserShop)
	result.BaseTable.Init(result)
	return result
}
