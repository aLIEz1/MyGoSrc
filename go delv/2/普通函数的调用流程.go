package main

func funcc(c int)  {
	println("c=",c)
}
func funcb(b int)  {
	funcc(b-1)
	println("b=",b)

}
func funca(a int)  {
	funcb(a-1)
	println("a=",a)
}
func main()  {
	funca(3)
	println("main")
}
//函数调用流程：先进后出
//函数递归：函数调用自己本身，利用了这个特点
