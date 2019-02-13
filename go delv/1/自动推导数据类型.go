package main

import "fmt"

func main() {
	var a int
	a = 10 //赋值，赋值前必须先声明变量，可以使用n次
	fmt.Println("a = ", a)
	b := 20                //:= ,自动推导数据类型，先声明变量b，再赋值，同一个变量名只能使用一次，用于初始化
	fmt.Println("b = ", b) //前面有b不能再b:= xxx 不能再新建一个b
}

//printf 格式化输出，把a的内容放在%d的输出
//println一段一段处理，自动加换行
