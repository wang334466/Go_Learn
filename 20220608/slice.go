package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	// 创建一个长度为3的空切片
	var slice2 []int
	slice2 = make([]int, 3)
	// 一行代码声明
	var slice3 []int = make([]int, 3)
	// 自定义寻找格式，声明slice是一个切片同时分配空间
	slice4 := make([]int, 3)
	slice5:=make([]int,0)
	//  %d 十进制表示，%v相应值的默认形式
	fmt.Printf("len=%d,slice1=%v\n", len(slice1), slice1)
	fmt.Printf("len=%d,slice2=%v\n", len(slice2), slice2)
	fmt.Printf("len=%d,slice3=%v\n", len(slice3), slice3)
	fmt.Printf("len=%d,slice4=%v\n", len(slice4), slice4)
	fmt.Printf("len=%d,slice4=%v\n", len(slice5), slice5)
}
