package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func WriteFile(path string)  {
	rand.Seed(time.Now().UnixNano())
	f,err:=os.Create(path)
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for i:=0;i<100 ;i++  {
		buf:=fmt.Sprintf("%d\n",rand.Intn(1024))
		fmt.Println("buf=",buf)
		n,err:=f.WriteString(buf)
		if err!=nil {
			fmt.Println(err)
			return
		}
		fmt.Println("n=",n)
	}

}
func main()  {
	path:="./mydemo.txt"
	WriteFile(path)
}
