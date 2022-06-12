package main

import "fmt"

// 定义一个结构体，多种类型数据的重新定义
type Book struct {
	title  string
	authon string
}


func changeBook(book Book) {
	book.authon = "666"

}

func changeBook2(book *Book) {
	book.authon = "777"
}
func main() {
	var book1 Book
	book1.title = "golang"
	book1.authon = "zhangsan"

	fmt.Printf("%v\n", book1)
	// 不会修改值
	changeBook(book1)

	fmt.Printf("%v\n", book1)
	// 传递指针，修改值
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)

}
