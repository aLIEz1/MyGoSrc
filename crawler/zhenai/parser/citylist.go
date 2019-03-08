package parser

import (
	"crawler/engine"
	"regexp"
)

//ParseCityList
const CityListRe  = `linkContent:"(?s:(.*?))",linkURL:"(http://m.zhenai.com/zhenghun/[0-9a-z]+)"`
func ParseCityList(contents []byte)engine.ParseResult{
	re := regexp.MustCompile(CityListRe)
	submatch := re.FindAllSubmatch(contents, -1)
	result:=engine.ParseResult{}
	for _, m := range submatch {
		result.Items=append(result.Items,string(m[1]))
		result.Requests=append(result.Requests,engine.Request{Url: string(m[2]),ParserFunc:engine.NilParser})
	}
	return result
}
