package main

import (
	"errors"
	"fmt"
)

// 高阶函数：接受其他的函数作为参数传入；或者把其他的函数作为结果返回
type operate func(x, y int) int

// 方案1。
func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

// 方案2。
type calculateFunc func(x int, y int) (int, error)

func genCalculator(op operate) calculateFunc {
	// 返回一个闭包函数
	// 闭包函数：在一个函数中存在对外来标识符的引用op(自由变量)
	// 闭包体现的是由「不确定」变为「确定」的一个过程，因为引用了「自由变量」而呈现出了一种不确定的状态（开放状态）
	// 也就是说闭包函数的内部逻辑并不是完整的，有一部分逻辑需要这个自由变量参与完成，而自由变量代表了什么在闭包函数被定义的时候却是未知的
	return func(x int, y int) (int, error) {
		// 自由变量op被捕获，此时闭包函数的状态由「不确定」变为「确定」
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	// 方案1。
	x, y := 12, 23
	op := func(x, y int) int {
		return x + y
	}
	result, err := calculate(x, y, op)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)
	// 函数类型属于引用类型
	result, err = calculate(x, y, nil)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)

	// 方案2。
	x, y = 56, 78
	add := genCalculator(op)
	result, err = add(x, y)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)

	x, y = 4, 5
	op = func(x, y int) int {
		return x * y
	}
	multiply := genCalculator(op)
	result, err = multiply(x, y)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)

}
