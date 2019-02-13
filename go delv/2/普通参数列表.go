package main

import "fmt"

//有参无返回值函数的定义，普通参数列表
//定义函数时，在函数名后面（）定义的参数叫形参
//参数传递是单向传递，只能由实参传递给形参
func MYfunc(a int){
	//a=111
	println("a= ",a)

}
func MYFUNC2(a int,b int)  {
	fmt.Printf("a=%d,b=%d",a,b )
}
func main(){
	//调用函数传递的叫实参
	MYfunc(668)
	MYFUNC2(667,556)

}