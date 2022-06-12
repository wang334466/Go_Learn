package main

import "fmt"

func defer1() {
	fmt.Println("defer1...")
}

func defer2() {
	fmt.Println("defer2...")
}

func defer3() {
	fmt.Println("defer3...")
}

func main() {
	// 栈，先进后出原则
	defer defer1()
	defer defer2()
	defer defer3()
}
