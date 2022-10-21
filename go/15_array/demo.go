package main

import "fmt"

func main() {
	//数组的定义1
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 12
	arr1[2] = 44
	arr1[1] = 21
	fmt.Println("第一种定义数组的方式", arr1)

	// 2
	var arr2 = [3]int{1, 12, 44}
	fmt.Println("第二种定义数组的方式", arr2)

	// 3
	var arr3 = [...]int{1, 12, 44}
	fmt.Println("第三种定义数组的方式", arr3)

}
