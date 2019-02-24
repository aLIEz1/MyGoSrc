package main

import (
	"fmt"
	"net"
)

func main()  {
	conn,err:=net.Dial("tcp","127.0.0.1:8000")
	if err!=nil {
		fmt.Println("err=",err)
		return
	}
	defer conn.Close()
	//buf:=make([]byte,2048)
	_,err1:=conn.Write([]byte("HelloIt'sMe"))//[]byte("可写字符串自动转换成字节")
	if err1!=nil {
		fmt.Println("err1=",err1)
		return
	}

}