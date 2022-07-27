/*
	slice切片的追加
*/

package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)

	fmt.Printf("numbers的长度为 = %d, numbers的cap = %d, numbers = %v\n", len(numbers), cap(numbers), numbers)

	//最加一个值为1
	numbers = append(numbers, 1)
	fmt.Printf("numbers的长度为 = %d, numbers的cap = %d, numbers = %v\n", len(numbers), cap(numbers), numbers)

	//最加一个值为2
	numbers = append(numbers, 2)
	fmt.Printf("numbers的长度为 = %d, numbers的cap = %d, numbers = %v\n", len(numbers), cap(numbers), numbers)

	//最加一个值为3，此时超过设置的cap=5，自动增加一倍的cap
	numbers = append(numbers, 1)
	fmt.Printf("numbers的长度为 = %d, numbers的cap = %d, numbers = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("------------------------------")

	// 可以不设置cap，默认和长度一致 （最常用）
	var numbers2 = make([]int, 3)
	fmt.Printf("numbers的长度为 = %d, numbers2的cap = %d, numbers2 = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 1)
	fmt.Printf("numbers的长度为 = %d, numbers2的cap = %d, numbers2 = %v\n", len(numbers2), cap(numbers2), numbers2)
}
