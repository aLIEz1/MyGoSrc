package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GETURL(url string) (result string) {
	resp,err1:=http.Get(url)
	if err1!=nil {
		fmt.Println(err1)
		return
	}
	buf:=make([]byte,4096)
	for {
		n,_:= resp.Body.Read(buf)
		if n == 0 {
			//fmt.Println("read err1=", err2)
			break
		}
		result += string(buf[:n])
	}
	defer resp.Body.Close()
	return
}
func SpiderJokePage(i int,page chan int)  {
	fmt.Println("working page",i,"...")
	url:="https://www.pengfue.com/index_"+strconv.Itoa(i)+".html"
	result:=GETURL(url)
	reg1:=regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if reg1==nil {
		fmt.Println("regexp.MustCompile err")
		return
		}
	Website:=reg1.FindAllStringSubmatch(result,-1)
	FILETITLE:=make([]string,0)
	FILECONTENT:=make([]string,0)
	for _,URLS:=range Website{
		SPIDERJOKE(URLS[1])
		title,content,_:=SPIDERJOKE(URLS[1])
		FILETITLE=append(FILETITLE,title)
		FILECONTENT=append(FILECONTENT,content)
		}
	WRITEFILE(FILETITLE,FILECONTENT,i)
	page<-i
}
func WRITEFILE(title,content []string,i int)  {
	FileName:=strconv.Itoa(i)+".txt"
	F,err3:=os.Create(FileName)
	if err3!=nil {
		fmt.Println("err3=",err3)
		return
	}
	n:=len(title)
	for i:=0;i<n;i++{
		//F.WriteString("Title:"+title[i]+"\n"+"Content:"+content[i])
		F.WriteString("Title:"+title[i]+"\n")
		F.WriteString("Content:"+content[i]+"\n")
		F.WriteString("\n===================\n")
	}
	defer F.Close()
}
func DoMyWorks1(startP,endP int)  {
	page:=make(chan int)
	for i:=startP;i<=endP ;i++  {
		go SpiderJokePage(i,page)

	}
	for i:=startP;i<=endP ;i++  {
		fmt.Println("page",<-page,"done")

	}
}
func SPIDERJOKE(URLS string) (title,content string,errs error) {
	result:=GETURL(URLS)
	reg1:=regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if reg1==nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	TI:=reg1.FindAllStringSubmatch(result,1)
	for _,data:=range TI{
		title=data[1]
		title=strings.Replace(title,"\t","",-1)
		break
	}
	reg2:=regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if reg2==nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	CO:=reg2.FindAllStringSubmatch(result,1)
	for _,data:=range CO{
		content=data[1]
		content=strings.Replace(content,"\t","",-1)
		content=strings.Replace(content,"\r","",-1)
		content=strings.Replace(content,"\n","",-1)
		content=strings.Replace(content,"<br />","",-1)
		content=strings.Replace(content,"&nbsp;","",-1)
		break
	}
	//joke<-URLS
return
}
func main()  {
	var startP,endP int
	fmt.Printf("Please input strat page:")
	fmt.Scan(&startP)
	fmt.Printf("Please input end page:")
	fmt.Scan(&endP)
	fmt.Printf("working %d--->%d...",startP,endP)
	DoMyWorks1(startP,endP)
}
