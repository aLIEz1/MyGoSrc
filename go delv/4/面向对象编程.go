package main

import "fmt"

type student struct {
	id int
	name string
	age int
} 
type STU student
//带有接收者的函数叫做方法
func (t STU)Printinfo()  {
	fmt.Println(t)
}
func (t *STU)Setinfo(n string,i int,a int)  {
	//fmt.Printf("请输入学生信息：\n")
	//fmt.Scan(t.id)
	//fmt.Scan(t.name)
	//fmt.Scan(t.age)
	t.name=n
	t.id=i
	t.age=a
}
func main()  {
	var s STU
	s.name="lily"
	s.id=111
	s.age=33
	(&s).Setinfo("mike",12,33)
	s.Printinfo()
}