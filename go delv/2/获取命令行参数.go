package main

import (
	"fmt"
	"os"
)//导入包必须使用，否则编译不过

func main()  {
	//接受用户传递的参数都是以字符串方式传递
	list:=os.Args
	n:=len(list)
	println("n=",n)
	for i,data :=range list{
		fmt.Printf("list[%d]=%s\n",i,data )
	} 
	
}
