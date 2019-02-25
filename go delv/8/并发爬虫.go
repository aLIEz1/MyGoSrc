package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)
func SpiderPage(i int,page chan int)  {
	fmt.Println("working page",i,"...")
	url:="http://tieba.baidu.com/f?kw=%E6%88%91%E8%A6%81%E5%BD%93%E5%AD%A6%E9%9C%B8&ie=utf-8&pn="+strconv.Itoa((i-1)*50)
	var result string
	resp,err1:=http.Get(url)
	if err1!=nil {
		fmt.Println(err1)
		return
	}
		buf:=make([]byte,4096)
		for{
			n,err2:=resp.Body.Read(buf)
			if n==0 {
				fmt.Println("read err=",err2)
				break
			}
			result+=string(buf[:n])
			FileName:=strconv.Itoa(i)+".html"
			F,err3:=os.Create(FileName)
			if err3!=nil {
				fmt.Println("err3=",err3)
				return
			}
			F.WriteString(result)
			defer F.Close()
		}
		defer resp.Body.Close()
		page<-i
}
func DoWorks(startP,endP int)  {
	page:=make(chan int)
	for i:=startP;i<=endP ;i++  {
		go SpiderPage(i,page)

	}
	for i:=startP;i<=endP ;i++  {
		fmt.Println("page",<-page,"done")

	}
}
func main()  {
	var startP,endP int
	fmt.Printf("Please input strat page:")
	fmt.Scan(&startP)
	fmt.Printf("Please input end page:")
	fmt.Scan(&endP)
	fmt.Printf("working %d--->%d...",startP,endP)
	DoWorks(startP,endP)
}