package main

import (
	"fmt"
)

func printArray(myArray []int) {
	// 引用传递
	// _表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value=", value)
	}
	myArray[0] = 100
}

func main() {
	// 动态数组 slice 切片
	myArray := []int{1, 2, 3, 4}
	// 打印myArray的数据类型
	fmt.Printf("myArray type is %T\n", myArray)
	printArray(myArray)

	fmt.Println("==========")
	for _, value := range myArray {
		fmt.Println("value=", value)
	}

}
