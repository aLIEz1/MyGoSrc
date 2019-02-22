package main

import "fmt"

func main()  {
	ch:=make(chan int)
	go func() {
		for i:=0;i<6 ;i++  {
			ch<-i
		}
		close(ch)//关闭channel     关闭后不可以再向channel里写入数据但是可以读取
	}()
	/*for{
		if num,ok:=<-ch;ok==true {
			fmt.Println(num)
		}else {
			break
		}
	}*/
	for num:=range ch{//通过range形式更加简洁
		fmt.Println(num)
	}
}
