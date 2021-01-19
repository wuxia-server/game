package Data

type DropItem struct {
	ItemId int
	Num    int
}

func (e *DropItem) ToJsonMap() map[string]interface{} {
	return map[string]interface{}{
		"item_id": e.ItemId,
		"num":     e.Num,
	}
}

func (e *DropItem) ToArray() []int {
	return []int{e.ItemId, e.Num}
}
