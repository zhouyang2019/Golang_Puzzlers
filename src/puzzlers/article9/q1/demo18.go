package main

func main() {
	// 示例1。
	// 键类型的值必须要支持判等操作，函数类型、字典类型、切片类型的值不支持判等操作
	//var badMap1 = map[[]int]int{} // 这里会引发编译错误。
	//_ = badMap1

	// 示例2。
	// 如果键的类型是接口类型的，那么键值的实际类型也不能是上述三种类型（函数、字典、切片）
	//var badMap2 = map[interface{}]int{
	//	"1":      1,
	//	[]int{2}: 2, // 编译器不会报错，但是runtime这里会引发panic。
	//	3:        3,
	//}
	//_ = badMap2

	// 示例3。
	// 如果键的类型是数组类型，还要确保该类型的「元素类型不是函数类型、字典类型、切片类型」
	//var badMap3 map[[1][]string]int // 这里会引发编译错误。
	//_ = badMap3

	// 示例4。
	//type BadKey1 struct {
	//	slice []string // 注意 切片类型
	//}
	//var badMap4 map[BadKey1]int // 这里会引发编译错误。
	//_ = badMap4

	// 示例5。
	//var badMap5 map[[1][2][3][]string]int // 这里会引发编译错误。
	//_ = badMap5

	// 示例6。
	//type BadKey2Field1 struct {
	//	slice []string
	//}
	//type BadKey2 struct {
	//	field BadKey2Field1
	//}
	//var badMap6 map[BadKey2]int // 这里会引发编译错误。
	//_ = badMap6

}
