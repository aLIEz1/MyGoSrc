package main

import "fmt"

type FuncType func()
func main()  {
	a:=10
	str:="mike"
	f1:= func() {
		println("a=",a)
		println("str=",str)
	}
	f1()
	var f2 FuncType  =f1
	f2()
	//定义匿名函数同时调用
	func (){
		println("a=",a)
		println("str=",str)
	}()
	func (i,j int){
		fmt.Printf("i=%d,j=%d\n",i,j )
	}(10,20)
	//匿名函数有参有返回值
	x,y:=func(i,j int )(max,min int){
		if i>j {
			max=i
			min=j
		}else {
			max=j
			min=i
		}
		return
	}(10,20)
	fmt.Printf("max=%d,min=%d\n",x,y )

}