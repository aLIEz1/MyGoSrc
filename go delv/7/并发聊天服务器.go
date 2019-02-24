package main

import (
	"fmt"
	"net"
)

var OnlineMap map[string]ChatInfo
var message =make(chan string)//定义管道没有分配空间不会有错误提示！！！！！！！定义的时候要分配好空间避免类似错误再次发生

type ChatInfo struct {
	C chan string//用户发送数据的管道
	Addr string//网络地址
	Name string//用户名
}

func HandleConnect(conn net.Conn)  {
	defer conn.Close()
	//OnlineMap=make(map[string]ChatInfo)
	Addr:=conn.RemoteAddr().String()
	cilent:=ChatInfo{make(chan string),Addr,Addr}
	OnlineMap[Addr]=cilent
	message<-MakeMsg(cilent,"Login")
	go WriteInfo(cilent,conn)
	Exitch:=make(chan bool)
	go TransmitMsg(cilent,conn,Exitch)//实现消息转发

	for{
		select {
		case <-Exitch:
			message<-MakeMsg(cilent,"Logout")
			delete(OnlineMap,cilent.Addr)
			return

		}

	}

}
func MakeMsg(Cilent ChatInfo,messagE string)(Msg string)  {
	Msg="["+Cilent.Addr+"]"+Cilent.Name+":"+messagE
	return
}
func Manager()  {
	OnlineMap=make(map[string]ChatInfo)//map初始化
	for{
		msg:=<-message//message没有赋值之前阻塞
		for _,cilent:=range OnlineMap{
			cilent.C<-msg
		}
	}
}
func WriteInfo(cilent ChatInfo,conn net.Conn)  {
	defer conn.Close()
	for msg:=range cilent.C{
		conn.Write([]byte(msg+"\n"))
		//if err!=nil {
		//	fmt.Println(err)
		//	return
		//}
	}

}
func TransmitMsg(cilent ChatInfo,conn net.Conn,Exitch chan bool)  {
	defer conn.Close()
	buf:=make([]byte,1024)
	for{
		n,err:=conn.Read(buf)
		if n==0 {//出错或者用户主动退出
			Exitch<-true
			fmt.Println("err",err)
			return
		}
		msg:=string(buf[:n-1])
		if msg=="who" {
			for _,data:=range OnlineMap{
				conn.Write([]byte("["+data.Addr+"]"+":"+data.Name+"\n"))
			}

		}else if msg=="rename" {
			conn.Write([]byte("Please input your new name:"))
			BUF:=make([]byte,1024)
			n,errs:=conn.Read(BUF)
			if errs!=nil {
				fmt.Println(errs)
			}
			cilent.Name=string(BUF[:n-1])
			OnlineMap[cilent.Addr]=cilent
			conn.Write([]byte("Rename Successfully"+"\n"))
		}else if msg=="exit" {
			message<-MakeMsg(cilent,"Logout")
			delete(OnlineMap,cilent.Addr)
			conn.Write([]byte("Exit Successfully"+"\n"))
			return
		}else {
			message<-MakeMsg(cilent,msg)
		}

	}

}
func main()  {
	listener,err:=net.Listen("tcp","127.0.0.1:8000")
	if err!=nil {
		fmt.Println("net.Listen err=",err)
		return
	}
	defer listener.Close()

	go Manager()
	for{
		conn,err:=listener.Accept()
		if err!=nil {
			fmt.Println("listener.Accept err=",err)
			continue
		}
		go HandleConnect(conn)

	}
}
