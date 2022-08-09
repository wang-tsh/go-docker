// Package vo @Title
// @Description  请填写文件描述
// @Author  wangtingshun  2022/8/9 15:29
// @Update  wangtingshun  2022/8/9 15:29
package vo

type OrderItem struct {
	column string // 需要排序的字段
	asc    bool   // 是否正序排列，默认true
}

func BuildAsc(column string) OrderItem {
	return OrderItem{column: column, asc: true}
}

func BuildDesc(column string) OrderItem {
	return OrderItem{column: column, asc: false}
}

func BuildAscs(columns ...string) []OrderItem {
	items := make([]OrderItem, 0)
	for _, val := range columns {
		items = append(items, BuildAsc(val))
	}
	return items
}

func BuildDescs(columns ...string) []OrderItem {
	items := make([]OrderItem, 0)
	for _, val := range columns {
		items = append(items, BuildDesc(val))
	}
	return items
}
