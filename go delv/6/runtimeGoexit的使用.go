package main

import (
	"fmt"
	"runtime"
)

func test()  {
	defer fmt.Println("cccccccccc")
	//return//终止此函数，不会终止上一级函数
	runtime.Goexit()//只打印aaaccc退出当前所在的协程Goexit（）//匿名函数是协程，终止了匿名函数
	fmt.Println("ddddddddddd")
}

func main()  {
	go func() {
		fmt.Println("aaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbb")
	}()
	for{

	}
}
