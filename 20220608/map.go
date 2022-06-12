package main

import (
	"fmt"
	// "strings"
)

func main() {
	// 默认为null
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")

	}
	// 声明2、直接创建
	myMap2 := make(map[string]string)
	myMap2["one"] = "java"
	myMap2["two"] = "c++"
	myMap2["three"] = "python"
	// map随机输出
	fmt.Println(myMap2)
	// 声明3：直接初始化
	myMap3 := map[int]string{
		1: "php",
		2: "c",
		3: "net",
	}
	fmt.Println(myMap3)

}
