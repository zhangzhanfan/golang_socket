/*
	slice的定义
*/

package main

import "fmt"

func main() {
	// 第一种：声明一个切片，并且初始化，默认值为1，2，3
	// slice1 := []int{1, 2, 3}

	//第二种：声明一个切片，开辟的空间为3，默认值为0
	// var slice1 []int
	// slice1 = make([]int, 3)

	//第三种：声明一个切片，同时开辟空间为3，默认值为0
	// var slice1 []int = make([]int, 3)

	//第四种：用 := 声明一个切片，同时go会根据make自动识别类型，开辟空间为3，默认值为0 （最常用的一种方法）
	slice1 := make([]int, 3)

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}
}
