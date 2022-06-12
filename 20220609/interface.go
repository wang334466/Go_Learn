/*
多态：
有一个父类（接口）
有子类（实现了父类全部接口方法）
父类类型的变量指针指向子类的具体数据变量
*/
package main

import "fmt"

// 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GerType() string  //获取动物的种类
}

// 具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GerType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}
func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) GerType() string {
	return "Dog"
}

// 接口的数据类型，父类指针
func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color=", animal.GetColor())
	fmt.Println("Kind=", animal.GerType())
}
func main() {
	cat := Cat{"Blue"}
	dog := Dog{"Yellow"}
	showAnimal(&cat)
	showAnimal(&dog)

}
