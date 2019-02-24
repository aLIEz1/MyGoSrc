package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func FileS(s string,conn net.Conn)  {
	File,err:=os.Open(s)
	if err!=nil {
		fmt.Println("os.Open err=",err)
		return
	}
	defer File.Close()
	con:=make([]byte,1024)
	for{
		_,err1:=File.Read(con)
		if err1!=nil {
			if err1==io.EOF {
				break
			}
			fmt.Println("File.Read err",err1)
	}
		_,err2:=conn.Write(con)
		if err2!=nil {
			fmt.Println("conn.Write err=",err2)
			return
		}


	}
	defer conn.Close()
}
func main()  {
	conn,err:=net.Dial("tcp","127.0.0.1:8000")
	if err!=nil {
		fmt.Println("net.Dial err=",err)
		return
	}
	defer conn.Close()
	s:=os.Args
	if len(s)!=2 {
		fmt.Println("usage:xxx filename")
		return
	}
	info,err1:=os.Stat(s[1])
	if err1!=nil {
		fmt.Println("os.Stat err=",err1)
	}
	_,err2:=conn.Write([]byte(info.Name()))
	if err2!=nil {
		fmt.Println("conn.Write err",err2)
	}
	//ch:=make(chan bool)
	//select {
	//case ok:=<-ch:
	//	if ok==false {
	//		return
	//	}else {
	//		File(s[2],conn)
	//	}
	//}
	buf:=make([]byte,1024)
	n,err3:=conn.Read(buf)
	if err3!=nil {
		fmt.Println(err3)
		return
	}
	if string(buf[:n])=="ok" {
		FileS(s[1],conn)
	}
}
