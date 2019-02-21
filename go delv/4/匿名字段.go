package main

import "fmt"

type Person struct {
	Name string
	Age int
	Sex byte
}
type Student struct {
	Person//只有类型没有名字	。匿名字段，继承了Person的成员
	Id int
	Addr string
}
type Student2 struct {
	*Person
	id int
	addr string
}
func main()  {
	var s2 Student= Student{Person{"mikke", 18, 'f'}, 1616, "bj"}
	fmt.Println("s2=",s2)
	s1:=Student{Person{"liming",12,'m'},1313,"jk"}
	fmt.Println("s1=",s1)
	s3:=Student{Person:Person{Name:"lily"},Id:1}//部分初始化
	fmt.Printf("s3=%+v\n",s3)
	s1.Name="yoyo"//使用.操作来操作成员，和指针类似
	fmt.Println("s1=",s1)
	s1.Person=Person{"go",16,'f'}//操作结构体，实现改多个成员变量
	fmt.Println("s1=",s1)
	//GO语言中的 默认原则（就近原则），如果能够在本作用域找到此成员，就操作此成员，如果没有找到，就找继承的字段
	s1.Person.Name="gogo"//显式调用
	fmt.Println("s1=",s1)
	s4:=Student2{&Person{"lili",20,'m'},1231,"az"}//结构体指针类型变量初始化方法1
	fmt.Printf("s4=%+v\n",s4)
	var s5 Student2//定义变量//方法2
	s5.Person=new(Person)//分配空间
	s5.Name="mk"
	s5.id=12123
	s5.addr="we"
	s5.Age=45
	s5.Sex='m'
	fmt.Printf("s5=%+v\n",s5)
}
