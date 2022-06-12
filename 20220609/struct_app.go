package main

import "fmt"

// Hero 如果类型大写，表示其他的包也能访问
type Hero struct {
	// 如果属性首字母大写，表示该属性是对外能够访问的，否则的话只能内部访问
	Name  string
	Ad    int
	Level int
}

// this是调用该方法的对象的一个副本（拷贝），引用传递
// 使用*，指针指向地址，可改变结果体内部的值
func (this *Hero) show() {
	fmt.Println("Name=", this.Name)
	fmt.Println("Ad=", this.Ad)
	fmt.Println("Level=", this.Level)
}
func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	hero := Hero{Name: "zhangsan", Ad: 100}
	hero.show()
	fmt.Println("====================")
	hero.SetName("lisi")
	hero.show()
}
