// 并行：多个cpu共同进行
// 并发：单个cpu，多个线程共同进行
// 进程：独享资源，进程崩溃影响较小，相互独立
// 线程：共享进程资源，线程崩溃，导致进程挂掉
// goroutine：协程、processor：处理器、thread：线程
package main

import (
	"fmt"
	"time"
)

// 子Goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine :i=%d\n", i)
		time.Sleep(1 * time.Second)

	}
}

// 主Goroutine
func main() {
	// 创建一个go程，去执行newTask()
	go newTask()
	// i := 0
	// for {
	// 	i++
	// 	fmt.Printf("main Goroutine i =%d\n", i)
	// 	time.Sleep(1 * time.Second)
	// }
	fmt.Println("main Goroutine finish")

}
