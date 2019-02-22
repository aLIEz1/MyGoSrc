package main

import (
	"fmt"
	"time"
)

func main()  {
	ch:=make(chan int,3)
	go func() {
		for i:=0;i<100 ;i++  {
			fmt.Println("Son=",i)
			ch<-i
		}
		//time.Sleep(7*time.Second)
	}()
	time.Sleep(7*time.Second)
	for i:=0;i<100 ;i++ {
		num:=<-ch
		fmt.Println("Main=",num)

	}

}
