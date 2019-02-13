package main

import "calc"

func main()  {
	a:=calc.Add(10,50)
	println("a=",a)
}//调用别的包的函数，这个包函数名字如果是小写，无法调用，要想让别人调用必须首字母大写
