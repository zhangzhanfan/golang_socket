package main

import "fmt"

func main() {
	str := "zhangzhanfan"
	for i, v := range str {
		fmt.Print(i)
		fmt.Printf("%c\t", v)
	}
}
