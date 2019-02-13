package main

import "fmt"

//...int类似这样的类型...type不定参数类型
//传递的实参可以是0个或者是多个
//不定参数只能放在形参中的“最后”一个参数
//固定参数一定要传参，不定参数可以传参可以不传参
//func myfunc03(args ...int,a int)  {
//
//}错误示范
func myfunc5(args ...int){
	println("len(args)=",len(args))//h获取用户传递参数的个数
	fmt.Printf("____________________________\n")
	for i,data :=range args {//返回两个值第一个是下标第二个是下标所对应的数
		fmt.Printf("args[%d]=%d\n",i,data )
	}

}
func main(){
myfunc5()
myfunc5(1)
myfunc5(2,30)
myfunc5(2,3,4,5,6)

}