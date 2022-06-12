package main

import (
	"fmt"
)

// 返回单个返回值
func foot1(a int, b int) int {
	fmt.Println("------foot1------")
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 123
	return c
}

// 返回多个返回值
func foot2(a int, b int) (int, int) {
	fmt.Println("------foot2------")
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	return 100, 200

}

// 形参带有名称的
func foot3(a int, b string) (r1 int, r2 int) {
	fmt.Println("------foot3------")
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	r1 = 1000
	r2 = 1100
	return
}

func main() {
	c := foot1(1, 2)
	fmt.Println(c)

	re1, re2 := foot2(400, 401)
	fmt.Println(re1, re2)

	ge1, ge2 := foot3(55, "abc")
	fmt.Println(ge1, ge2)

}
