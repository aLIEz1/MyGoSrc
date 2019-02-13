package main

import "fmt"

func main() {
	var ch byte
	var str string
	//字符是单引号，字符串是双引号，字符往往只有一个字符，转义字符除外
	//字符串有一个或者多个字符组成
	//字符串往往都是隐藏了一个结束符'\0'
	ch = 'a'
	str = "a"
	fmt.Println("ch=", ch)
	fmt.Println("str=", str)
	str = "hello go"
	fmt.Printf("str[0]=%c,str[1]=%c\n", str[0], str[1])
}
