package main
//*p指向实参的内存，是实参数组a的指针
func modify(p *[5]int) {
	(*p)[0] = 666
	for _, data := range *p {
		println("modify", data)
	}
}
func main()  {
	a:= [5]int{1,2,3,4,5}
	modify(&a)
	for _,data:=range a{
		println("main",data)
	}
}