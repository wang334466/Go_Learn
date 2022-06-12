package main

import (
	"fmt"
	/*
		array 固定长度的数组，使用前必须确定数组的长度
		1，go中的数组是值类型
		2，如果go中的数组作为函数的参数的话，实际是传递一份数组的拷贝。而不是数组的指针
		3，数组的长度也是type的一部分，[10]int 和[20]int是不一样的

		slice是可变长的，定义一个slice变量之后，不需要为它的容量而担心，你随时可以往slice里面加数据
		var :=[]string{}
		v=append(v,"hello")
		注意:slice和array的写法很容易混淆
		v ：=[2]string{"str1","str2"} //这个是数组
		m ：=[]string{"str1","str2"} //这个是slice
		写法上array有长度slice没有长度

		slice是一个指针而不是值。
		指针比值小很多，因此，我们将slice作为函数传递比数组传递更有性能。
		slice是一个指针，它指向的是一个array结构，他们都有len和cap。

		map结构
		map结构跟php的array几乎一模一样，是一个key-value的hash结构，key可以是除了func类型，array，slice，map类型之外的类型。
	*/)

func main() {
	// 创建切片，3=数量，5=容量，当大于这个容量则底层成倍扩展容量。
	numbers := make([]int, 3, 5)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	// 添加数据
	numbers = append(numbers, 1)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 2)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	fmt.Println("==================")
	var number2 = make([]int, 3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(number2), cap(number2), number2)

	s := []int{1, 2, 3}
	// 数据截取
	s1 := s[0:2]
	fmt.Println(s1)
	// 数据可修改
	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1)
	// 创建空切片
	s2 := make([]int, 3)
	// 拷贝数据
	copy(s2, s)
	fmt.Println(s2)
}
