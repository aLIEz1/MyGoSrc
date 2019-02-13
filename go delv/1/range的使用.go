package main

import "fmt"

func main() {
	str := "HelloIAmYourFather"
	for i, data := range str {
		fmt.Printf("str[%d]=%c\n", i, data)
	}
}
