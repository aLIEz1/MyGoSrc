package main

func main()  {
	//a:=10
	var p *int
	//p=&a
	p=new(int)//go语言 具有自动的内存回收不用释放空间无需担心生命周期和怎样将其删除
	*p=6678
	println("*p=",*p)
	q:=new(int)
	*q=777
	println("*q=",*q)
}