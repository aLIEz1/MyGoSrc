package main

import (
	"math/rand"
	"time"
)

func main()  {
	var a [20]int
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<20 ;i++  {
		a[i]=rand.Intn(20)
	}
	for i:=0;i<4 ;i++  {
		for j:=0;j<4-i ;j++  {
			if a[j]>a[j+1] {
				a[j],a[j+1]=a[j+1],a[j]
			}
		}
	}
	for _,data:=range a {
		println("a=", data)
		//println("a=",a)
	}


}