package main

import "fmt"

//回调函数，函数参数是函数类型，这个就是回调函数
//多态，多种形态，调用同一个接口，不同的表现，可以实现加减乘除
//先有想法再实现功能
type FuncType func(int,int)int

func add(a,b int) int {
	return a+b
}
func minus(a,b int)int  {
	return a-b
}
func multi(a,b int)int  {
	return a*b
}
func divi(a,b int)int  {
	return a/b
}
func Calc(a int ,b int ,test FuncType)(resuult int)  {
	println("Calc")
	resuult=test(a,b)
	return

}
func main()  {
	a:=Calc(1,10,add)
	b:=Calc(50,10,minus)
	c:=Calc(5,8,multi)
	d:=Calc(50,2,divi)
	//println("result=",a)
	fmt.Printf("1+10=%d,50-10=%d,5*8=%d,50/2=%d\n",a,b,c,d )
	
}