package main

import "fmt"

func main() {
	var a int
	fmt.Printf("请输入a：")
	//阻塞等到用户输入
	fmt.Scanf("%d", &a)
	fmt.Scan(&a)
	fmt.Println("a= ", a)
}
