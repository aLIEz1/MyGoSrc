package main

import (
	"fmt"
	"net"
)

func main()  {
	listener,err:=net.Listen("tcp","127.0.0.1:8000")//监听是否有请求
	if err!=nil {
		fmt.Println("err=",err)
		return
	}
	defer listener.Close()
	conn,err1:=listener.Accept()//接受请求
	if err1!=nil {
		fmt.Println("err1=",err1)
		return
	}
	con:=make([]byte,2048)
	n,err2:=conn.Read(con)//读取请求内容
	if err2!=nil {
		fmt.Println("err2=",err2)
		return
	}
	fmt.Println("con=",string(con[:n]))
	defer conn.Close()
}
