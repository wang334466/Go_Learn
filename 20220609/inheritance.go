package main

import "fmt"

// 声明父类
type Human struct {
	Name string
	sex  string
}

// 父类方法
func (this *Human) eat() {
	fmt.Println("Human eat...")
}

// 父类方法
func (this *Human) walk() {
	fmt.Println("Human walk ...")
}

// 声明子类，继承父类
type superHuman struct {
	Human
	level int
}

// 新创建子类方法
func (this *superHuman) fly() {
	fmt.Println("superHuman fly ...")
}

// 重定义子类方法walk（）
func (this *superHuman) walk() {
	fmt.Println("superHuman walk ...")
}

// 输出函数
func (this *superHuman) print() {
	fmt.Println("Name=", this.Name)
	fmt.Println("sex=", this.sex)
	fmt.Println("level=", this.level)
}
func main() {
	// 初始化方式一，父类对象 h
	h := Human{"zhangsna", "female"}
	h.eat()
	h.walk()
	// 初始化方式二，定义一个子类对象 s
	var s superHuman
	s.Name = "lisi"
	s.sex = "male"
	s.level = 88
	// 子类方法
	fmt.Println("=======子类=======")
	s.print()

	s.eat()
	s.walk()
	s.fly()
}
