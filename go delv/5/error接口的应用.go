package main

import (
	"errors"
	"fmt"
)

func MyDiv(a,b int)(result int ,err error)  {
	err=nil
	if b==0 {
		err=errors.New("分母不能为0")
	}else {
		result=a/b
	}
	return

}
func main()  {
	r,err:=MyDiv(12,0)
	if err!=nil {
		fmt.Println("error!",err)
	}else {
		fmt.Println("result=",r)

	}

}
