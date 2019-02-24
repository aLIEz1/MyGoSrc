package main

import (
	"fmt"
	"io"
	"net"
)

func RecvKeyboard(K *string,conn net.Conn)  {
	defer conn.Close()
	for{
		fmt.Scan(K)
		_,err:=conn.Write([]byte(*K))
		if err!=nil {
			if err==io.EOF {
				break
			}
			fmt.Println("err=",err)
			return
		}
	}

}
func RecvServer(conn net.Conn)  {
	con:=make([]byte,2048)
	for{
		n,err:=conn.Read(con)
		if err!=nil {
			//if nil==io.EOF {
			//	break
			//}
			fmt.Println("err=",err)
			return
		}
		fmt.Println(string(con[:n]))
	}

}
func main()  {
	conn,err:=net.Dial("tcp","127.0.0.1:8000")
	if err!=nil {
		fmt.Println("err=",err)
		return
	}
	defer conn.Close()
		var keyboard string
		go RecvKeyboard(&keyboard,conn)
		RecvServer(conn)
}