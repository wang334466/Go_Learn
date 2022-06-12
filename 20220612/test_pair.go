package main

import "fmt"

func main() {
	// 变量，type value 类型不变
	// pair<statictype:string,value:"abcls">
	var a string
	a = "abcls"
	// pair<statictype:string,value:"abcls">
	var allType interface{}
	allType = a
	str, _ := allType.(string)
	fmt.Println(str)
}
