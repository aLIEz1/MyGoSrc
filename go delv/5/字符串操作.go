package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Contains("helloworld","hello"))
	fmt.Println(strings.Contains("helloworld","abf"))//判断字符里有没有指定的字符，有则输出true没有则输出false
	s:=[]string{"mike","love","lily"}
	tmp:=strings.Join(s,"+")//用+号将切片里的字符连接起来
	fmt.Println("join=",tmp)
	fmt.Println(strings.Index("helloworld","world"))
	fmt.Println(strings.Index("helloworld","make"))//查找指定字符串在母字符串中的位置，有则输出位置，没有则输出-1
	tmp="mike*like*lily"
	s2:=strings.Split(tmp,"*")//以*分割字符串tmp
	fmt.Println("split=",s2)
	tmp=strings.Trim("aaaacccccaaaa","a")//去掉两头的a
	fmt.Println("trim=",tmp)
	s3:=strings.Fields("are you ok i am fine fuck you")//去掉空格以空格分割字符串转换成切片
	fmt.Println("field=",s3)

}
