package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	jsBuf:=`{
	"isok": true,
	"name": "heima",
	"price": 66.66666,
	"subject": [
		"Go",
		"C#",
		"C++",
		"Python"
	]
}
`
	m:=make(map[string]interface{},4)
	err:=json.Unmarshal([]byte(jsBuf),&m)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("map=",m)
	//map里的数据只可以通过类型断言反推类型，再保存赋值
	for key,value:=range m{
		switch value.(type) {
		case string:
			fmt.Printf("m[%s]string====>%v\n",key,value)
		case bool:
			fmt.Printf("m[%s]bool====>%v\n",key,value)
		case float64:
			fmt.Printf("m[%s]float64====>%v\n",key,value)
		case []interface{}:
			fmt.Printf("m[%s][]interface{}====>%v\n",key,value)//subject用了万能指针空切片[]interface{}



			
		}
	}
}
