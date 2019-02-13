package main

import "fmt"

func main() {
	//iota常量自动生成器，每隔一行自动累加一
	//iota给常量赋值使用
	//iota遇到const，重置为0
	//iota可以只写一个iota
	//如果是同一行值都一样
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)
	const d = iota
	fmt.Println("d=", d)
	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	fmt.Printf("i=%d,j1=%d,j2=%d,j3=%d,k=%d\n", i, j1, j2, j3, k)
}
