package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

// go run demo4.go demo4_lib.go
func main() {
	flag.Parse()
	hello(name)
}
