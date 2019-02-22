package main

import (
	"fmt"
	"time"
)

func main()  {
	ch:=make(chan string)
	go func() {
		defer fmt.Println("子协程调用完毕")
		for i:=0;i<3;i++{
			fmt.Println("i=",i)
			time.Sleep(time.Second)
		}
		ch<-"我是子协程，我调用完毕"//调用完毕后传参，解锁
	}()
	str:=<-ch//ch没有数据会阻塞在这一步简称锁
	fmt.Println(str)
	defer fmt.Println("主协程执行完毕")
}
