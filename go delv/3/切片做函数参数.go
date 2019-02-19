package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int)  {
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<len(s) ;i++  {
		s[i]=rand.Intn(100)
	}



}
func BubbleSort(s []int)  {
	for i:=0;i<len(s)-1 ;i++  {
		for j:=0;j<len(s)-1-i ;j++  {
			if s[j]>s[j+1] {
				s[j],s[j+1]=s[j+1],s[j]
			}
		}
	}
}
func main()  {
	n:=10
	s:=make([]int,n)
	InitData(s)
	fmt.Println("排序前s=",s)
	BubbleSort(s)
	fmt.Println("排序后s=",s)
}