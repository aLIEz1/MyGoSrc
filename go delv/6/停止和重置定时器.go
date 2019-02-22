package main

import (
	"fmt"
	"time"
)

func main()  {
	timer:=time.NewTimer(3*time.Second)
	timer.Reset(19*time.Second)//将定时器重置为19秒
	go func() {
		<-timer.C
		fmt.Println("Son Over")
	}()
	timer.Stop()//停止定时器,此时什么也不会打印，因为主协程3秒的时间早已经将timer关闭了
	for{

	}
}
