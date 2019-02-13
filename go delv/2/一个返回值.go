package main
//无参有返回值,通过return 中断函数
func myfunc01()  int{
	return  666
}
func myfunc02()(result int)  {//最常用写法，注意有两个括号
	result=66666
	return
}
func main()  {
	var a int  =myfunc01()
	println("a=",a)
	b:=myfunc01()
	println("b=",b)
	c:=myfunc02()
	println("c=",c)
}