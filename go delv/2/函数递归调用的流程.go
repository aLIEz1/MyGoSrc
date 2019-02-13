package main

func test(a int)  {
	if a==1 {
		println("a=",a)
		return
	}
	test(a-1)
	println("a=",a)
}
func main()  {
	test(3)
	println("main")
}