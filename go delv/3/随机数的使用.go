package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//设置种子：只需要一次
	//如果种子参数一样，每次运行程序产生的随机数都一样
	rand.Seed(time.Now().UnixNano())//引入时间参数，时间每次都不一样，产生的随机数每次也都不一样
	for i:=0;i<5 ;i++  {
		//println("rand=",rand.Int())
		//println("rand2=",rand.Intn(100))//Intn(填写随机数的范围，最大值)
		fmt.Printf("rand[%d]=%d\n",i,rand.Intn(100) )
	}
	//产生随机数

}