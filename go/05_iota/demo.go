package main

import "fmt"

// iota, 自增加1，常量计数器，计算同一个const中的常量个数   只能用于常量
func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)

	const (
		a1 = iota   //iota 0
		b1          //iota 1
		c1          //iota 2
		d1 = "haha" //iota 3
		e1          //iota 4
		f1 = 100    //iota 5
		g1          //iota 6
		h1 = iota   //iota 7
		i1          //iota 8
	)
	fmt.Println(a1, b1, c1, d1, e1, f1, g1, h1, i1)
}
