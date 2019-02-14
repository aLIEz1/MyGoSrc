package main
//s数组做函数参数，是值传递【【实参数组每个元素会给形参数组拷贝一份
func modify(a [5]int)  {
	a[0]=666
	for _,data:=range a{
		println("modify",data)
	}

}
func main()  {
	a:= [5]int{1,2,3,4,5}
	modify(a)
	for _,data:=range a{
		println("main",data)
	}

}