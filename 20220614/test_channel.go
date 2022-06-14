package main

import "fmt"

func main() {
	// 创建一个channel 无缓冲
	c := make(chan int)
	// 匿名函数
	go func() {
		defer fmt.Println("goroutine is fished")
		fmt.Println("goroutine is running")
		// 将666发送给c
		c <- 666
	}()
	// 声明将c值发送给num
	num := <-c

	fmt.Println("num=", num)
	fmt.Println("main goroutine is fished")

}
