package main

func main()  {
	a:=10
	str:="mike"
	func(){
		//闭包是以引用方捕获外部变量（里面改了外面也改）
		a=666
		str="go"
		println("inside:a=",a)
		println("inside:str=",str)

	}()
	println("outside:a=",a)
	println("outside:str=",str)
}