package main

import "fmt"

// go语言中，函数是一等公民，函数类型也是一等数据类型
// 所以函数不但可以用于封装代码、分割功能、解耦逻辑，还可以化身为普通的值，在其他函数间传递、赋予变量、做类型判断和转换等等
// 对函数类型来说，它是一种对一组输入、输出进行模块化的重要工具，比接口更加轻巧灵活

// 声明一个函数类型Printer,注意类型声明的名称右边是func关键字
type Printer func(contents string) (n int, err error)

// 函数签名：函数的参数列表和结果列表的统称
func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}
// 只要两个函数的参数列表和结果列表中的元素顺序及其类型是一致的，我们就可以说它们是实现了同一个函数类型的函数
func printToStd2(c string) (int, error) {
	return fmt.Println(c)
}
func printToStd3(i int) (int, error) {
	return fmt.Println(i)
}

func main() {
	var p Printer
	p = printToStd
	p("something")

	p = printToStd2
	p("method2")

	printToStd3(123)

}
