package main

import (
	"fmt"
	"runtime"
)

func main()  {
	go func() {
		for i:=0;i<5;i++  {
			fmt.Println("go")
		}
	}()
	for i:=0;i<2; i++ {
		runtime.Gosched()//让别的协程先执行，别的执行完了再来执行此协程，接力棒---》
		fmt.Println("hello")
	}
}
