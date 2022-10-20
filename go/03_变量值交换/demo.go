package main

import "fmt"

func main() {
	var a int = 100
	var b int = 240
	a, b = b, a
	fmt.Print(a, b)
}
