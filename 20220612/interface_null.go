package main

import (
	"fmt"
)
// interface[]万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called ...")
	fmt.Println(arg)
    // 给interface 提供"类型断言"
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type,value = ", value)
		fmt.Printf("arg type is %T\n", value)
	}

}

type Book struct {
	auth string
}

func main() {
	book := Book{"golang"}
	myFunc(book)
	myFunc(100)
	myFunc(3.14)
	myFunc("abc")
}
