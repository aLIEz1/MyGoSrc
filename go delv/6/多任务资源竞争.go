package main

import (
	"fmt"
	"time"
)

func Printer(s string)  {
	for _,data:=range s{
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
}
func person1()  {
	Printer("HELLO")
}
func person2()  {
	Printer("WORLD")
}
func main()  {
	go person1()//person1和person2竞争资源，导致打印错误信息
	go person2()
	for{

	}
}
