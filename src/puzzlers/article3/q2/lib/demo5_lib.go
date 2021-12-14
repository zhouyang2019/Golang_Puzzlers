package lib

import "fmt"

// 首字母为大写的程序实体才可以被当前包外的代码引用，否则它就只能被当前包内的其他代码引用
// 包级公开
func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 包级私有
func hello2(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 模块级私有，puzzlers/article3/q4/lib/internal/internal.go
// internal代码包中声明的公开程序实体仅能被该代码包的直接父级包及其子包中的代码引用