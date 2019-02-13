package main

func main()  {
	a:=10
	b:=20
	defer func(a,b int) {
		println("inside:a=",a)
		println("inside:b",b)
	}(a,b)//执行到此先把参数传递过去只是没有调用
	a=111
	b=222
	println("outside:a=",a)
	println("outside:b",b)

}