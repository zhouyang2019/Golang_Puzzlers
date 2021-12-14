package main

import "fmt"

type DbItem struct {
	Id int16
	Cnt int32
}

func combineItemWithSlice(itemList []DbItem, id int16, cnt int32)  {
	item := DbItem{id, cnt}
	itemList[0] = item
	fmt.Println("combineItemWithSlice itemList values %v \n", itemList)
}

func combineItemWithArray(itemList [5]DbItem, id int16, cnt int32) {
	item := DbItem{id, cnt}
	itemList[0] = item
	fmt.Println("combineItemWithArray itemList values %v \n", itemList)
}

func appendItem(itemList []DbItem, id int16, cnt int32) {
	item := DbItem{id, cnt}
	itemList = append(itemList, item)
	fmt.Println("appendItem itemList values %v \n", itemList)
}

func appendItem2(itemList []DbItem, id int16, cnt int32) []DbItem {
	item := DbItem{id, cnt}
	itemList = append(itemList, item)
	fmt.Println("appendItem itemList values %v \n", itemList)
	return itemList
}

func main()  {
	// 切片是引用类型，参数为引用副本
	itemList := make([]DbItem, 5)
	combineItemWithSlice(itemList, 1, 2)
	fmt.Println("main itemList values %v \n", itemList)

	// 数组是值类型，参数为值副本
	itemList2 := [5]DbItem{{Id: 3, Cnt: 4}}
	combineItemWithArray(itemList2, 10, 20)
	fmt.Println("main itemList2 values %v \n", itemList2)

	// 因为参数为引用副本，所以即使传的slice是一个引用类型，也是传的它的一个copy
	itemList3 := make([]DbItem, 5)
	appendItem(itemList3, 100, 200)
	fmt.Println("appendItem itemList3 values %v \n", itemList3)
	itemList3 = appendItem2(itemList3, 100, 200)
	fmt.Println("appendItem itemList3 values %v \n", itemList3)


}
