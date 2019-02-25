package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func DoMyWorks(url string ,i int)(err error)  {
	resp,err1:=http.Get(url)
	if err1!=nil {
		err=err1
	}
	buf:=make([]byte,4096)
	var result string
	for{
		n,err2:=resp.Body.Read(buf)
		if n==0 {
			err=err2
			break
		}
		result+=string(buf[:n])
	}
	reg1:=regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if reg1==nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	Website:=reg1.FindAllStringSubmatch(result,-1)
	for _,data:=range Website {
		title,content,err:=SpideJoke(data[1])
		if err!=nil {
			continue
		}
		fmt.Println("title:",title)
		fmt.Println("content:",content)
	}
return
}
func SpideJoke(Urls string) (title string,content string,err error) {
	resp,err1:=http.Get(Urls)
	if err1!=nil {
		err=err1
	}
	buf:=make([]byte,4096)
	var result string
	for{
		n,err2:=resp.Body.Read(buf)
		if n==0 {
			err=err2
			break
		}
		result+=string(buf[:n])
	}
	reg1:=regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if reg1==nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	reg1.FindAllStringSubmatch(title,1)
	reg2:=regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if reg2==nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	reg2.FindAllStringSubmatch(content,-1)
	//defer resp.Body.Close()
return 
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
		url:="https://www.pengfue.com/index_"+strconv.Itoa(i)+".html"
		go DoMyWorks(url,i)
		err:=DoMyWorks(url,i)
		if err!=nil {
			fmt.Println("DoWork err=",err)
			continue
		}
	}
}
