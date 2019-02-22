package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func WriteFile(path string)  {
	rand.Seed(time.Now().UnixNano())
	file,err:=os.Create(path)
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i:=0;i<100 ;i++  {
		con:=fmt.Sprintf("%d\n",rand.Intn(100))//将打印结果转换成字符串存储在con中
		_,err:=file.WriteString(con)
		if err!=nil {
			fmt.Println(err)
			return
		}
	}
}
func ReadFileLine(path string)  {
	file,err:=os.Open(path)
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//con:=make([]byte,4096)
	//n,err1:=file.Read(con)//n表示读取的文件的长度
	//if err1!=nil&&err!=io.EOF{//文件出错同时没有到结尾
	//	fmt.Println(err1)
	//	return
	//}
	//fmt.Println(string(con[:n]))
	r:=bufio.NewReader(file)//创建缓冲区，把内容放在缓冲区     利用bufio实现按行读取
	for{
		con,err:=r.ReadBytes('\n')//遇到‘\n’结束读取
		if err!=nil {
			if err==io.EOF {//到文件末尾时跳出循环
				break
			}else {
				fmt.Println(err)
			}
		}
		fmt.Printf("con=#%s#\n",string(con))//会把'\n”读取进去
	}

}
func main()  {
	path:="./go delv/5/mydemo.txt"
	//WriteFile(path)
	ReadFileLine(path)
}
