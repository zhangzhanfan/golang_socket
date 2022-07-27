/*
	slice切片的截取
*/

package main

import "fmt"

func main() {
	// s1和s是指向同一个地址空间
	s := []int{1, 2, 3}
	s1 := s[0:2] //截取第0和第1位元素
	fmt.Println(s1)
	// s1修改，s也跟着修改
	s1[0] = 101
	fmt.Println(s)
	fmt.Println(s1)

	// 将s中的值拷贝到s2
	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println(s2)
	s2[0] = 102
	fmt.Println(s2)
	fmt.Println(s)
	fmt.Println(s1)
}
