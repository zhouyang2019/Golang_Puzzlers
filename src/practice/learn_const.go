package main

import "fmt"

func main()  {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of area : %d", area)   //原文出自【易百教程】，商业转载请联系作者获得授权，非商业请保留原文链接：https://www.yiibai.com/go/go_constants.html#article-start

	//LENGTH = 20 // cannot assign to const

}
