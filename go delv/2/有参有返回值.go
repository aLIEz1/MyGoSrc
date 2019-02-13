package main

import "fmt"

func maxandmin(a,b int)(max,min int)  {
	if a>b {
		max=a
		min=b
	} else {//else必须写在if的花括号后面不能另起一行
		max=b
		min=a
	}
	return
}
func main()  {
	max,min:=maxandmin(10,20)
	fmt.Printf("max=%d,min=%d\n",max,min )
}