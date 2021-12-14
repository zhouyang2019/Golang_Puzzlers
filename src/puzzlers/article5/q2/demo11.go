package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	fmt.Printf("The element is %q.\n", container[1]) // The element is one.
	fmt.Printf("The element is %v.\n", container[1]) // The element is one.
	fmt.Printf("The element is %+v.\n", container[1]) // The element is one.
	// 不同类型的可重命名变量
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1]) // The element is "one".
	fmt.Printf("The element is %v.\n", container[1]) // The element is one.
	fmt.Printf("The element is %+v.\n", container[1]) // The element is one.

	container2 := make(map[[2]int16]string)
	arr1 := [2]int16{1, 2}
	arr2 := [2]int16{1, 3}
	container2[arr1] = "s1"
	container2[arr2] = "s2"
	fmt.Printf("container2 is %q \n", container2)
	fmt.Printf("container2[arr1] %q \n", container2[arr1])
}
