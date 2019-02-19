package main

import "fmt"

type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}
func main()  {
	var s1 Student  =Student{1,"mike",'m',20,"beijing"}//顺序初始化，每个成员都必须初始化
	fmt.Println("s1=",s1)
	s2:=Student{name:"jack",addr:"tianjin"}//指定成员初始化，没有初始化的成员赋值为0
	fmt.Println("s2=",s2)
	var p1 *Student  =&Student{1,"mike",'m',20,"beijing"}//不要忘记加*
	p2:=&Student{name:"jack",addr:"tianjin"}
	fmt.Println("p1=",p1)//GO语言支持直接打印p1不加*也可打印指针p1指向的内存
	fmt.Println("p2=",p2)
	var s3 Student
	s3.id=123//操作成员使用.操作
	s3.name="lm"
	s3.sex='m'
	s3.age=20
	s3.addr="zj"
	fmt.Println("s3=",s3)
	var s4 Student//第一种方法：先定义一个结构体变量，再定义一个结构体指针指向它
	var p3 *Student  =&s4//指针有合法指向后再操作成员
	p3.name="mk"//p3.name="mk"  (*p3).id=1212这两种写法完全一样，都可以
	(*p3).id=1212
	p3.sex='f'
	p3.age=12
	p3.addr="am"
	fmt.Println("p3=",p3)
	p4:=new(Student)//第二种方法：通过new来申请一个结构体指针
	p4.sex='f'
	p4.addr="jn"
	fmt.Println("p4=",p4)
	test(&s1)//引用传递，形参可以改实参
	fmt.Println("s1=",s1)
}
func test(p *Student)  {
	p.id=6666
	fmt.Println("s=",p)
}
