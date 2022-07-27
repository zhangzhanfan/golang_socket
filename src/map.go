/*
	map定义
*/

package main

import "fmt"

func main() {
	// 第一种：
	var map1 map[string]string
	if map1 == nil {
		fmt.Println("map1是一个空map")
	}
	map1 = make(map[string]string, 3)
	map1["one"] = "java"
	map1["two"] = "C语言"
	map1["three"] = "C++"
	fmt.Println(map1)

	// 第二种 不赋值常用
	map2 := make(map[int]string)
	map2[1] = "中国"
	map2[2] = "美国"
	map2[3] = "缨国"
	fmt.Println(map2)

	//第三种 赋值常用
	map3 := map[int]string{
		1: "懂得",
		2: "往往",
		3: "啊啊",
	}
	fmt.Println(map3)

}
