package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents,err:=ioutil.ReadFile("citylist.html")
	if err!=nil {
		panic(err)

	}
	fmt.Printf("%s\n",contents)
}
