package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Name string `json:"name"`//不会输出到字段
	Subjects []string `json:"subjects"`//转换成首字母小写输出到字段，但是注意原来的必须首字母大写
	IsOk bool `json:"isok"`//以字符串方式输出
	Price float64`json:"price"`
}
func main()  {
	jsBuf:=`{
	"subjects": [
		"Go",
		"C++",
		"Python",
		"Ruby"
	],
	"IsOk": true,
	"Price": 23.34
}
`


	var tmp IT
	err:=json.Unmarshal([]byte(jsBuf),&tmp)//解码，第二个参数必须要用取地址，否则改不了，只是复制了一份
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("tmp=%+v",tmp)
}