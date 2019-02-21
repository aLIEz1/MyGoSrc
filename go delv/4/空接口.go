package main

import "fmt"

type mk struct {
	name string
	id int
}
func WanNeng(arg ...interface{})  {//接收任意数量任意类型的参数

}
func main()  {
	var i interface{}=1//空接口  万能类型  保存任意类型的值
	fmt.Println("i=",i)
	i="abgksahdkj"
	fmt.Println("i=",i)
	s:=make([]interface{},3)
	s[0]=123
	s[1]="make"
	s[2]=mk{"mike",1313}
	fmt.Println("s=",s)

}