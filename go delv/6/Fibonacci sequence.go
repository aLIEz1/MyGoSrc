package main

import "fmt"

func Filbonacci(ch chan<- int,quit <-chan bool)  {
	x,y:=1,1
	for{
		select {
		case ch<-x:x,y=y,x+y
		case <-quit: return


		}
	}
}
func main()  {
	ch:=make(chan int)
	quit:=make(chan bool)
	go func() {
		for i:=0;i<16 ;i++  {
			fmt.Println(<-ch)
		}
		quit<-true
	}()
	Filbonacci(ch,quit)
}
