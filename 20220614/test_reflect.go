package main

import (
	"fmt"
	"reflect"
)

func PrintNum(arg interface{}) {
	// 反射机制
	fmt.Println("Type=", reflect.TypeOf(arg))
	fmt.Println("value=", reflect.ValueOf(arg))
}
func main() {
	var num float64 = 3.1425
	PrintNum(num)
}
