package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
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
func SpiderTUPage(i int,page chan int)  {
	fmt.Println("working page",i,"...")
	url:="https://gratisography.com/page/"+strconv.Itoa(i)+"/"
	os.Mkdir(strconv.Itoa(i),0666)
	result:=GETURL(url)
	reg1:=regexp.MustCompile(`<a href="(?:(.*?))" rel="bookmark"`)
	if reg1==nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	Website:=reg1.FindAllStringSubmatch(result,-1)
	LENWEB:=len(Website)
	for s:=0;s<LENWEB ;s++  {
		SPIDERTU(Website[s][1],s,i)
	}
	page<-i
}
func DoMyWorks1(startP,endP int)  {
	page:=make(chan int)
	for i:=startP;i<=endP ;i++  {
		go SpiderTUPage(i,page)

	}
	for i:=startP;i<=endP ;i++  {
		fmt.Println("page",<-page,"done")

	}
}
func SPIDERTU(URLS string,s int,i int) {
	result:=GETURL(URLS)
	reg1:=regexp.MustCompile(`<div class="download-buttons"><a target="_blank" href="(?s:(.*?))"`)
	if reg1==nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	DL:=reg1.FindAllStringSubmatch(result,1)
	buf:=GETTU(DL[0][1])
	StoreFile(i,s,buf)
}
func StoreFile(i,s int,buf []byte)  {
	ioutil.WriteFile(fmt.Sprintf("./%d/%d.jpg",i,s),buf,0644)
}
func GETTU(downloadLink string) ( data2 []byte) {
	resp,err1:=http.Get(downloadLink)
	if err1!=nil||resp.StatusCode!=200 {
		fmt.Println("Download failed",err1)
		return nil
	}
	data2,_=ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
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
