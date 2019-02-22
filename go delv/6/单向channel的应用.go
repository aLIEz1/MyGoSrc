package main

import "fmt"
//此通道只可以写不可以读
func producer(in chan <- int)  {
	for i:=0;i<10 ;i++  {
		in<-i*i*i
	}
	close(in)
}
//此通道只可以读不可以写
func consmer(out <-chan int )  {
	for num:=range out{
		fmt.Println("num=",num)
	}

}
//生产一个消费一个
func main()  {
	ch:=make(chan int)
	go producer(ch)//生产者，只读
	consmer(ch)//消费者，只写
}
