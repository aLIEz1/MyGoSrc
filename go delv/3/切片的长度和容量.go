package main

func main()  {
	a:=[...]int{1,2,3,4,0,0}
	s:=a[0:4:5]
	for _,data:=range s{
		println("s",data)
	}
	println("len(s)=",len(s))//长度4-0
	println("cap(s)=",cap(s))//容量5-0
}