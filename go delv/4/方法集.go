package main

import "fmt"

type Person struct {
	name string
	age int
	sex byte
}

func (p Person)SetInfoValue()  {
	fmt.Println("SetInfoValue")
}
func (p *Person)SetInfoPointer()  {
	fmt.Println("SetInfoPointer")
}
func main()  {
	p:=&Person{"mike",18,'m'}
	p.SetInfoPointer()
	p.SetInfoValue()
	//*P的方法集包含全部receiver P+*P的方法
	//GO语言会转换p和（*p）
	p2:=Person{"mike",18,'m'}
	p2.SetInfoValue()
	p2.SetInfoPointer()//普通变量也可以直接调用receiver为指针的函数，GO语言实现了自动类型转换
}