package main

import (
	"fmt"
	"io"
	"os"
)

var x int16 = 66

func main() {
	var err error
	// 被声明的变量必须是多个，并且其中有一个是新的变量。这时我们才可以说对其中的就变量进行了重声明
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n") // 这里对`err`进行了重声明。
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%d byte(s) were written.\n", n)

	var s string = "str1"
	fmt.Println("s: %s, x: %d", s, x)
	s, x := "str2", 2 // 当前变量覆盖外层变量
	fmt.Println("s: %s, x: %d", s, x)


}
