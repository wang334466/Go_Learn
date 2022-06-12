package main

import (
	"20220607/bb/lib1"
	// 匿名导包，导入不使用报错，因此采用_来匿名
	// _ "20220607/bb/lib2"
	// 导入所有的包，不需要.来引入
	. "20220607/bb/lib2"
)

func main() {
	lib1.Lib1Test()
	// lib2.Lib2Test()
	Lib2Test()
}
