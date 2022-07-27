/*
	反射2
*/
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

func (this User) Call() {
	fmt.Println("调用了Call方法")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{Id: 1, Name: "小黄", Age: 13}
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	// 获取input的类型
	inputType := reflect.TypeOf(input)
	fmt.Println(inputType.Name())

	// 获取input的值
	inputValue := reflect.ValueOf(input)
	fmt.Println(inputValue)

	fmt.Println("-------------------------")

	//通过type获取里面的字段
	// 1 获取interface的reflect.type,通过type得到numfield,进行遍历
	// 2 得到每个field，数据类型
	// 3 通过field 有一个interface方法等到 对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s -- %v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s  %v\n", m.Name, m.Type)
	}
}
