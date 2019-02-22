package main

import (
	"fmt"
	"time"
)

var ch  =make(chan int)
func Printer(s string)  {
	for _,data:=range s{
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
}
func person1()  {
	Printer("HELLO")
	ch<-666//向管道发送数据
}
func person2()  {
	<-ch//从管道接收数据，没有数据会{堵塞}在这一步，直到有数据为止
	Printer("WORLD")
}
func main()  {
	go person1()//person1和person2竞争资源，导致打印错误信息
	go person2()
	for{

	}
}