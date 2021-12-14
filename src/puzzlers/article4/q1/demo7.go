package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string                                                   // [1]
	flag.StringVar(&name, "name", "everyone", "The greeting object.") // [2]

	// 方式1。类型推断（编译器自动解释表达式类型）： 声明变量name的同时赋值，但是声明中并没有显示指定变量类型。
	//var name = flag.String("name", "everyone", "The greeting object.")

	// 方式2。短变量声明（类型推断+语法躺）：只能在函数体内部使用，变量重声明
	//name := flag.String("name", "everyone", "The greeting object.")

	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)

	// 适用于方式1和方式2。
	//fmt.Printf("Hello, %v!\n", *name)
}
