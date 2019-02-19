package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreatNum(p *int)  {
	var num  int
	rand.Seed(time.Now().UnixNano())
	for {
		num=rand.Intn(10000)
		if num>=1000 {
			break
		}
	}
	*p=num
}
func TakeNum(s []int,p *int){
	s[0]=*p/1000
	s[1]=(*p%1000)/100
	s[2]=(*p%100)/10
	s[3]=*p%10
}
func OnGame(RandSlice []int)  {
	num:=0
	var InputSlice  = make([]int,4)
	for {
		for {
			fmt.Printf("请输入一个四位数：")
			fmt.Scan(&num)
			if num > 999 && num < 10000 {
				break
			} else {
				println("输入的数不符合要求")
			}
		}
		TakeNum(InputSlice, &num)
		n:=0
		for i:=0;i<4 ;i++  {
			if InputSlice[i]>RandSlice[i] {
				fmt.Printf("第%d位大了\n",i+1 )
			} else if InputSlice[i]<RandSlice[i] {
				fmt.Printf("第%d位小了\n",i+1 )
			} else {
				fmt.Printf("第%d位猜对了\n",i+1 )
				n++
			}
		}
		if n==4 {
			fmt.Printf("\n全部猜对了!!!\n")
			break
		}
	}
}
func main()  {
	RandNum:=0
	CreatNum(&RandNum)
	//println("rand number=",RandNum)
	var RandSlice = make([]int, 4)
	TakeNum(RandSlice,&RandNum)
	//fmt.Println("RandSlice",RandSlice)
	OnGame(RandSlice)
}