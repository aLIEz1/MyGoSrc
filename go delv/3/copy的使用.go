package main

import "fmt"

func main()  {
	s1:=[]int{1,2,3}//不是直接替换，而是替换指定位置
	s2:=[]int{6,6,6,6,6,6}
	copy(s2,s1)
	fmt.Println("s2=",s2)
}
