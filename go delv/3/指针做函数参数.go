package main

func main()  {
	a,b:=10,20
	swap(&a,&b)
	println("a=",a)
	println("b=",b)
}
func swap(a,b *int)  {
	//var temp *int
	//temp=a
	//a=b
	//b=temp可以这样写，但是没必要
	println("a=",*a)
	println("b=",*b)

	*a,*b=*b,*a//是把值互换而不是地址互换
	println("a=",*a)
	println("b=",*b)

}