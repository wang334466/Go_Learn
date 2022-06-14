package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

// 定义父类方法
func (this User) Call() {
	fmt.Println("user is called ...")
	fmt.Println("%v\n", this)
}
func (this User) Eat() {
	fmt.Println("Human eat...")
}

func DoMethon(input interface{}) {
	// 反射类型
	inputType := reflect.TypeOf(input)
	fmt.Println("input type=", inputType)
	// 反射有效值
	inputValue := reflect.ValueOf(input)
	fmt.Println("input value=", inputValue)
	// 输出类型type中的具体数据类型，值
	// 通过type获取里面的字段
	// 1.获取interface的reflect的reflect.type通过type获取numfield，进行遍历
	// 2.得到每一个field，数据类型
	// 3.通过field有一个interface方法得到对应的value
	// numfield返回结构体的属性个数
	for i := 0; i < inputType.NumField(); i++ {
		// field（）返回第几个属性值
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s:%v=%v\n", field.Name, field.Type, value)
	}
	fmt.Println(inputType.NumMethod())
	// 通过type 获取里面的方法，调用
	// 没有调用
	for j := 0; j < inputType.NumMethod(); j++ {
		m := inputType.Method(j)
		fmt.Println(123)
		fmt.Printf("%s:%v\n", m.Name, m.Type)
	}
}

func main() {
	user := User{12, "zhangsan", 25}
	DoMethon(user)
}
