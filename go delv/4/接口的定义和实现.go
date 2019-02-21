package main

import "fmt"

type Humer interface {
	sayHi()
}
type Student struct {
	name string
	id int
}
type Teacher struct {
	addr string
	group string
}
type PR struct {
	lth int
}

func (s *Student)sayHi()  {
	fmt.Println("stu")
}
func (s *Teacher)sayHi()  {
fmt.Println("teacher")
}
func (s *PR)sayHi()  {
fmt.Println("PR")
}
func WhoSayHi(i Humer)  {
	i.sayHi()
}
func main()  {
	var i Humer
	s:=&Student{"mike",1818}
	i=s
	i.sayHi()
	t:=&Teacher{"nj","go"}
	i=t
	i.sayHi()
	p:=&PR{12}
	i=p
	i.sayHi()
	//WhoSayHi(s)
	//WhoSayHi(t)
	//WhoSayHi(p)//调用同一函数不同表现，多态
	x:=make([]Humer,3)//创建切片来调用
	x[0]=s
	x[1]=t
	x[2]=p
	for _,i:=range x{
		i.sayHi()
	}
	
}
