package main

import (
	"fmt"
	"time"
)

func main()  {
	timer:=time.NewTimer(3*time.Second)//时间到了只会往timer.C里面写一次数据，不会写第二次；；；设置的时间是3秒
	fmt.Println(time.Now())
	t:=<-timer.C
	fmt.Println(t)
	//三大延时实现方式；；；上面是第一种
	//time.Sleep(3*time.Second)//第二种
	//<-time.After(3*time.Second)//第三种
}
