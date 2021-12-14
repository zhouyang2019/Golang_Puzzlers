package main

import (
	"flag"
	"puzzlers/article3/q4/lib"
	//in "puzzlers/article3/q4/lib/internal" // 此行无法通过编译。模块级私有：internal代码包让一些程序实体仅仅能被当前模块中的其他代码引用。
	//"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
	//in.Hello(os.Stdout, name)
}
