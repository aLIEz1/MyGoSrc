package main

import (
	"fmt"
	"time"
)

func sonTask()  {
	for{
		fmt.Println("sonTask")
		time.Sleep(time.Second)
	}
}
func main()  {
	go sonTask()//创建新的协程子协程和主协程竞争资源  go关键字，创建goroutine
	for{
		fmt.Println("mainTask")
		time.Sleep(time.Second)
	}
}
