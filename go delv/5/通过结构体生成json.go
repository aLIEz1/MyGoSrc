package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Name string `json:"-"`//不会输出到字段
	Subjects []string `json:"subjects"`//转换成首字母小写输出到字段，但是注意原来的必须首字母大写
	IsOk bool `json:",string"`//以字符串方式输出
	Price float64
}
func main()  {
	s:=IT{"heima",[]string{"Go","C++","Python","Ruby"},true,23.34}
	js,err:=json.MarshalIndent(s,"","	")//格式化打印
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("json=",string(js))
}
