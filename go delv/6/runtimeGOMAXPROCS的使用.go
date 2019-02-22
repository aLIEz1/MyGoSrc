package main

import (
	"fmt"
	"runtime"
)

func main()  {
	runtime.GOMAXPROCS(4)//设置以多少核来运行此程序
	for{
		go fmt.Print(1)
		fmt.Print(0)
	}
}