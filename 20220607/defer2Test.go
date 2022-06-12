package main

import "fmt"

func deferfunc() int {
	fmt.Println("defer..")
	return 0
}

func returnfunc() int {
	fmt.Println("return...")
	return 0
}

func retAndDer() int {
	defer deferfunc()
	return returnfunc()
}

// return与defer共存的情况下，return先返回
func main() {
	retAndDer()
}
