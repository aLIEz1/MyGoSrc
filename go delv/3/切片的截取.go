package main

func main()  {
	a:=[]int{1,2,3,4,5,6,7,8,9}
	for _,data:=range a{
		println(data)
	}
	println("===========")
	s1:=a[2:5]
	for _,data:=range s1{
		println(data)
	}
	println("===========")
	s2:=a[3:]
	for _,data:=range s2{
		println(data)
	}
	println("===========")
	s3:=a[:6]
	for _,data:=range s3{
		println(data)
	}
	println("===========")
	s4:=a[3:6]
	s4[0]=666
	for _,data:=range s4{
		println(data)
	}
	println("===========")
	for _,data:=range a{
		println(data)
	}
	println("===========")//切片可以看作指向数组的特殊指针，切片变，数组跟着变


}
