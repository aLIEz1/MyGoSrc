package main

import (
	"fmt"
	"net/http"
)

func main()  {
	resp,err:=http.Get("http://www.baidu.com")
	if err!=nil {
		fmt.Println("http.Get err",err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	buf:=make([]byte,4096)
	var tmp string
	for{
		n,err:=resp.Body.Read(buf)
		if n==0 {
			fmt.Println(err)
			break
		}
		tmp+=string(buf[:n])
	}
	fmt.Println(tmp)
}
