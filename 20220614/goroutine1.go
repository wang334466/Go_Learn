package main

import (
	"fmt"
	// "runtime"
	"time"
)

func main() {
	/*
		// 定义一个匿名函数
		go func() {
			defer fmt.Println("A.defer")
			// return
			// 子进程
			func() {
				defer fmt.Println("B.defer")
				// 退出当前的goroutine
				runtime.Goexit()
				fmt.Println("B")

			}()
			fmt.Println("A")
		}()
	*/

	// 匿名函数带参数
	go func(a int, b int) bool {
		fmt.Println("a=", a, "b=", b)
		return true
	}(10, 20)

	// 死循环
	for {
		time.Sleep(1 * time.Second)
	}

}
