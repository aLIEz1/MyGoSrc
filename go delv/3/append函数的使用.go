package main

import "fmt"

func main()  {
	s1:=[]int{}
	s1=append(s1,2,3,4,5,6,7)//append扩容容量特点：如果超过原来的容量，通常以两倍容量扩容
	fmt.Println("s1=",s1)
	println("len",len(s1))
	println("cap",cap(s1))
}