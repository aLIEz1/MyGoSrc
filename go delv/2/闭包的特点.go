package main

func test()int  {
	//函数被调用时，x被分配空间，初始化为0，调用完毕，x自动释放空间
	var x int//没有初始化，值为0
	x++
	return x*x
}
func test02()func()int  {
	var x int
	return func() int {
		x++
		return  x*x
	}
}
func main()  {
	//返回值为一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数，f来调用闭包函数
	//它不关心这些捕获的变量和常量是否已经超出了作用域
	//所以只有闭包还在使用它，这些变量就还是会存在
	f2:=test02()
	println(f2())
	println(f2())
	println(f2())
	println(f2())
	println(f2())
	println(f2())


}