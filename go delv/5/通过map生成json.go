package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	m:=make(map[string]interface{},4)
	m["name"]="heima"
	m["subject"]=[]string{"Go","C#","C++","Python"}
	m["isok"]=true
	m["price"]=66.66666
	js,err:=json.MarshalIndent(m,"","	")
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("map---->json",string(js))
	
}
