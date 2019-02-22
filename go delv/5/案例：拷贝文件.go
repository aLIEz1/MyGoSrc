package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Cmd := os.Args
	if len(Cmd) != 3 {
		fmt.Println("用法：SrcFile DstFile")
		return
	}
	SrcFile:=Cmd[1]
	DstFile:=Cmd[2]
	if SrcFile == DstFile{
		fmt.Println("源文件和目的文件不能同名")
		return
	}
	SF,errS:=os.Open(SrcFile)
	if errS!=nil {
		fmt.Println(errS)
		return
	}
	DF,errD:=os.Create(DstFile)
	if errD!=nil {
		fmt.Println(errD)
		return
	}
	defer SF.Close()
	defer DF.Close()
	con:=make([]byte,4*1024)//创建一个4k的缓冲区
	for{//一边读一边写，读多少写多少
		n,err:=SF.Read(con)//源文件读取，n是文件的长度
		if err!=nil {
			if err==io.EOF {
				break
			}else {
				fmt.Println(err)
			}
		}
		_,errW:=DF.Write(con[:n])////源文件写入  Con[:n]中的n是指读取的长度
		if errW!=nil {
			fmt.Println(errW)
			return
		}else {
			fmt.Println("拷贝成功！")
		}
	}

}
