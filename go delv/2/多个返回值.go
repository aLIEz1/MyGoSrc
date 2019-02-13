package main

import "fmt"

func myfunc011()(a,b,c int)  {
	a,b,c=111,222,333
	return
}
func main()  {
	a,b,c:=myfunc011()
	fmt.Printf("a=%d,b=%d,c=%d",a,b,c )
}