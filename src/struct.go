/*
	结构体
*/

package main

import "fmt"

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook(book *Book) {
	book.auth = "李四"
}

func main() {
	var book1 Book
	book1.title = "一本数额"
	book1.auth = "张三"
	fmt.Printf("%v\n", book1)

	changeBook(&book1)
	fmt.Printf("%v\n", book1)
}
