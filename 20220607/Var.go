package main

import "fmt"

func main() {
	// 正常声明变量
	var a int = 100
	fmt.Println(a)
	// %T输出数据类型
	fmt.Printf("Type of a=%T\n", a)
	// 默认值为0
	var b int
	fmt.Println(b)
	// 省略类型定义
	var c = "abc"
	fmt.Println(c)
	// 自动匹配类型
	// 重点：不可用于函数外
	d := 200
	fmt.Println(d)
	fmt.Printf("Type of d=%T\n", d)

	var (
		dc="string"
		dd=100
	)
	fmt.Println(dc,dd)

}
