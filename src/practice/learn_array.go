package main

import "fmt"

func main() {
	var arr1, arr2 [2]int16
	arr1 = [2]int16{1, 2}
	arr2 = [2]int16{1, 2}
	fmt.Printf("arr1 == arr2 ? %t \n", arr1 == arr2)

	var slice1, slice2 []int16
	slice1 = []int16{1}
	slice2 = []int16{1}
	fmt.Printf("slice1: %v \n", slice1)
	fmt.Printf("slice2: %v \n", slice2)
	//fmt.Println("slice1 == slice2 ? %t", slice1 == slice2) // operator == not defined on []int16

}
