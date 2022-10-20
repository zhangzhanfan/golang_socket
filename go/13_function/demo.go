package main

import "fmt"

// 函数也是一种类型,本案例可以证明函数是引用类型。
func main() {
	fmt.Printf("%T\n", f1)
	var f5 func(int, int)
	f5 = f1
	f5(1, 2)
	fmt.Println(f5)
	fmt.Println(f1)
}

func f1(a, b int) {
	fmt.Println(a, b)
}
