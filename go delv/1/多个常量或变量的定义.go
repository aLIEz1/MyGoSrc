package main

import "fmt"

func main() {
	//var a int
	//var b float64
	// var (
	// 	a int
	// 	b float64
	// )
	var ( //可以自动推导类型
		a = 1
		b = 2.34
	)
	//a, b = 11, 3.14
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	// const i int = 10
	// const j float64 = 3.14
	const (
		i = 10
		j = 3.14
	)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
}
