package main

func add(a int,b int)int  {
	return a+b
}
func minus(a int,b int)int  {
	return a-b
}

type FuncType func(int, int)int//没有函数名字没有{}定义在main函数之前
func main()  {
	var result int
	result=add(1,3)
	println("result=",result)
	var test FuncType  =minus//声明一个函数类型的变量，变量名字叫test
	result=test(10,20)//等价于minus(10,20)
	println("result2=",result)

}