package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func FileR(name string,coon net.Conn)  {
	Name:="./"+name
	File,err:=os.Create(Name)
	if err!=nil {
		fmt.Println("os.Create err=",err)
		return
	}
	defer File.Close()
	buf:=make([]byte,1024)
	for{
		_,err:=coon.Read(buf)
		if err!=nil {
			if err==io.EOF {
				break
			}
			fmt.Println(err)
			return
		}
		_,err1:=File.Write(buf)
		if err1!=nil {
			fmt.Println(err1)
			return
		}
	}
}
func main()  {
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err!=nil{
		fmt.Println("net.Listen err",err)
		return

	}
	defer listener.Close()
	conn,err1:=listener.Accept()
	if err1!=nil{
		fmt.Println("listener.Accept err",err1)
		return
	}
	con:=make([]byte,1024)
	n,err2:=conn.Read(con)
	if err2!=nil{
		fmt.Println("conn.Read err",err2)
		return
	}
	_,err3:=conn.Write([]byte("ok"))
	if err3!=nil{
		fmt.Println("conn.Write err",err3)
		return
	}
	FileR(string(con[:n]),conn)
	fmt.Println(time.Now(),"接收成功")
}
