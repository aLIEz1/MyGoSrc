package main

import "fmt"

type Person struct {
	name string
	sex byte
	age int
}
type Student struct {
	Person
	id int
	addr string
}

func (p *Person)PrintInfo()  {
	fmt.Println("person",p)
}
func (p *Student)PrintInfo()  {
	fmt.Println("student",p)

}
func main()  {
	s:=Student{Person{"mike",'m',14},1313,"gd"}
	s.PrintInfo()//s继承了p 的方法可以调用p专有的函数
	//如果存在同名函数，遵循就近原则，现在本作用域找，没有就找继承的方法
	s.Person.PrintInfo()//显式调用
	sPrint:=s.PrintInfo//方法值，调用函数时无需再传递接收者。隐藏了接收者
	sPrint()


}