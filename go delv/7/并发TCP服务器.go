package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func HandleConn(conn net.Conn)  {

	adr:=conn.RemoteAddr().String()
	fmt.Println(time.Now(),adr,"Connected successfully")
	for{
		con:=make([]byte,2048)
		n,err:=conn.Read(con)
		if err!=nil {
			fmt.Println("err=",err)
			return
		}
		fmt.Printf("%v[%s]---->%s",time.Now(),adr,string(con[:n]))
		if string(con[:n])=="exit" {
			fmt.Println(time.Now(),adr,"Exit successfully")
			return
		}else {
			_,err2:=conn.Write([]byte(strings.ToUpper(string(con[:n]))))
			if err2!=nil {
				fmt.Println("err2=",err2)
				return
			}
		}
		}
	defer conn.Close()
}
func main()  {
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err!=nil {
		fmt.Println("err=",err)
		return
	}
	defer listener.Close()
	for{
		conn,err:=listener.Accept()
		if err!=nil {
			fmt.Println("err=",err)
			return
		}
		go HandleConn(conn)
	}
}