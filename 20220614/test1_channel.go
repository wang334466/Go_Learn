package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建有缓存的channel
	c := make(chan int, 3)
	fmt.Println("len(c)=", len(c), ", cap(c)=", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		// 循环将i发送给c
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("子go进程正在运行，发送的元素为", i, "len(c)=", len(c), "cap(c)=", cap(c))

		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		// 从c中接收值，并赋值给num
		num := <-c
		fmt.Println("num=", num)
	}
	fmt.Println("main 结束")

}
