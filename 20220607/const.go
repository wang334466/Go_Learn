package main

import (
	"fmt"
)

// 声明枚举类型
const (
	// iota每一行增加1，从默认0开始
	a = iota * 10
	b
)

func main() {
	// 声明静态常量，常量不可以更改
	const lengh = 10
	fmt.Println("length=", lengh)
	// lengh =lengh+10
	fmt.Println("a=", a, "b=", b)
}
