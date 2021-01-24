package StaticTable

import (
	"github.com/team-zf/framework/Table"
	"github.com/team-zf/framework/logger"
	"github.com/team-zf/framework/utils"
)

type ShopDetail struct {
	Id             int         `ST:"PK"`                // 唯一ID
	GoodsBank      int         `ST:"goods_bank"`        // 商品库
	GoodsRank      int         `ST:"goods_rank"`        // 商品位置 (同一商品位置可能有多个数据取一条)
	GoodsDropId    int         `ST:"goods_drop"`        // 商品掉落ID
	GoodsRankProb  int         `ST:"rank_goods_weight"` // 商品位置里多物品权重
	GoodsCost      *Table.Map  `ST:"goods_cost"`        // 商品售价 (物品ID, 数量)
	UnlockCondIds  *Table.List `ST:"cond_unlock"`       // 商品解锁条件
	BuyCondIds     *Table.List `ST:"cond_buy"`          // 商品购买条件
	Discount       int         `ST:"discount"`          // 折扣类型 (0、不打折 1、售价打折 2、全新售价打折（重新定义售价消耗id和数量）)
	DiscountParams *Table.List `ST:"discount_para"`     // 折扣参数
	DiscountCanbuy int         `ST:"discount_canbuy"`   // 折扣次数
	NormalCanbuy   int         `ST:"normal_canbuy"`     // 正常购买次数 (正常购买次数限制不包含折扣次数)
}

var (
	_ShopDetailList []*ShopDetail
)

func init() {
	filePath := "./JSON/wx_shop_detail.json"
	rows, err := Table.LoadTable(filePath, &ShopDetail{})
	if err != nil {
		panic(err)
	}

	arr := make([]*ShopDetail, 0)
	for _, row := range rows {
		arr = append(arr, row.(*ShopDetail))
	}
	_ShopDetailList = arr
	logger.Notice("载入数据表[%s]", filePath)
}

func GetShopDetail(detailId int) (result *ShopDetail) {
	for _, row := range _ShopDetailList {
		if row.Id == detailId {
			newrow := utils.ReflectNew(row)
			result = newrow.(*ShopDetail)
			break
		}
	}
	return
}

func GetShopDetailList(goodsBank int) (result []*ShopDetail) {
	result = make([]*ShopDetail, 0)
	for _, detail := range _ShopDetailList {
		if detail.GoodsBank == goodsBank {
			result = append(result, detail)
		}
	}
	return
}
