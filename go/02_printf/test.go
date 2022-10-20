package main

import "fmt"

func main() {
	num := 100
	isFlag := true
	floatNum := 3.14
	fmt.Printf("num：%d, num的类型是：%T, num的内存地址：%p, isFlag的值为：%t\n", num, num, &num, isFlag)
	fmt.Printf("floatNum的值为：%f,floatNum的类型为：%T\n", floatNum, floatNum)
	fmt.Printf("floatNum的值为：%.1f,floatNum的类型为：%T\n", floatNum, floatNum)
	fmt.Println("str := \"c:\\pprof\\main.exe\"")
}
