package main

import "fmt"

func myfunc(tmp ...int)  {
	for _,data:=range tmp{
		fmt.Printf("data=%d\n",data )

	}

}
func Tets(args ...int)  {

	myfunc(args...)//全部元素传递给myfunc
	myfunc(args[:2]...)//args0~2不包括2传递过去
	myfunc(args[2:]...)//从args[2]开始包括本身，把后面所有元素传递过去
}
func main()  {
	Tets(1,2,3,4)
	//myfunc()
}