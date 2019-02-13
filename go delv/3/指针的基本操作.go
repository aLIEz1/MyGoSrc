package main

func main()  {
	var a int = 20
	println("&a=",&a)
	println("a=",a)
	var p *int  =&a//指针默认值nil没有null常量   不支持->访问变量地址，只支持.操作
	*p=666
	println("a=",a)
}
