package main

import "fmt"

//interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println(arg)
	// 判断传入的interface{}具体， 即 "类型断言" 机制
	value, ok := arg.(string) //判断arg是否为string类型
	if !ok {
		fmt.Println("arg不是 string 类型")
	} else {
		fmt.Println(value)
		fmt.Printf("%T\n", value)
	}
}

type Book struct {
	title string
}

func main() {
	book := Book{title: "地理"}
	myFunc(book)
	myFunc(100)
	myFunc(3.17)
	myFunc("abcd")
}
