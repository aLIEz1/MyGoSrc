package main

import "fmt"

func DEL(m map[int]string)  {
	delete(m,2)
}
func main()  {
	var m1 map[int]string //map只有len没有cap
	m2:=make(map[int]string,10)//只是指定了容量，但是没有赋值长度没有变还是0
	m2[1]="go"
	m2[2]="lang"
	m2[3]="C++"//map键值唯一，不能重复
	fmt.Println("m1=",m1)
	fmt.Println("m2=",m2)//map是无序的，无法决定返回顺序，每次打印结果顺序可能不同
	fmt.Println("len=",len(m2))
	m3:=map[int]string{1:"mike",2:"go",3:"C++"}
	m3[1]="jack"
	m3[4]="C#"//map底层自动扩容，和append类似
	fmt.Println("m3=",m3)
	for key,value:=range m3{//遍历map
		fmt.Printf("m3[%d]=%s\n",key,value )
	}
	_,ok:=m3[999]
	if ok==true {//判断是否存在这个key值
		fmt.Println("存在")
	}else {
		fmt.Println("不存在")
	}
	delete(m3,1)//删除key为1的内容
	DEL(m3)//map做函数参数是引用传递
}
