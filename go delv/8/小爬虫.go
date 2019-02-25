package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func DoWork(url string)(result string,err error)  {
	resp,err1:=http.Get(url)
	if err1!=nil {
		err=err1
	}
	defer resp.Body.Close()
	buf:=make([]byte,4096)
	for{
		n,err2:=resp.Body.Read(buf)
		if n==0 {
			fmt.Println("read err=",err2)
			break
		}
		result+=string(buf[:n])
	}
	return


}
func writeFiles(result string,i int)  {
	FileName:=strconv.Itoa(i)+".html"
	F,err:=os.Create(FileName)
	if err!=nil {
		fmt.Println("err=",err)
		return
	}
	F.WriteString(result)
}
func main()  {
	var startP,endP int
	fmt.Printf("Please input strat page:")
	fmt.Scan(&startP)
	fmt.Printf("Please input end page:")
	fmt.Scan(&endP)
	fmt.Println("working %d--->%d...",startP,endP)
	for i:=startP;i<=endP ;i++  {
		fmt.Println("working page",i,"...")
		url:="http://haodoo.net/?M=hd&P=100-"+strconv.Itoa(i)
		result,err:=DoWork(url)
		if err!=nil {
			fmt.Println("DoWork err=",err)
			continue
		}
		go writeFiles(result,i)
	}
}