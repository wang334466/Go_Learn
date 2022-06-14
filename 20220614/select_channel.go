package main

import "fmt"

func ff(c, quit chan int) {
	x, y := 1, 1
	for {
		// 多通道监控
		select {
		// c是否可写，将x发送给c
		case c <- x:
			x = y
			y = x + y
		// quit是否可读
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}
func main() {
	c := make(chan int)
	quit := make(chan int)
	// 子进程
	go func() {
		for i := 0; i < 10; i++ {
			// 从c中读，读出的数据输出
			fmt.Println(<-c)
			// 将i发送给c
			// c<-i
			// fmt.Println(c)
		}
		quit <- 0
	}()

	ff(c, quit)
}
