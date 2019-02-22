package main

import (
	"fmt"
	"time"
)
//主协程退出，子协程自动退出，如果主协程没有执行直接关闭了，子协程不能执行
func main()  {
	i:=0
	go func() {
		fmt.Println("son i=",i)
		i++
		time.Sleep(time.Second)
	}()
	for{
		fmt.Println("main i=",i)
		i++
		time.Sleep(time.Second)
		if i==4 {
			break
		}
	}
}
