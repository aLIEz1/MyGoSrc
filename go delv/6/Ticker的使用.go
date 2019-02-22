package main

import (
	"fmt"
	"time"
)

func main()  {
	ticker:=time.NewTicker(time.Second)//每隔一秒传一次数据，直到stop
	i:=0
	for{
		<-ticker.C
		fmt.Println("i=",i)
		i++
		if i==5 {
			ticker.Stop()
			break
		}
	}
}
