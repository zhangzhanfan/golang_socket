package main

import "fmt"

func main() {

	//定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("go goroutine 结束")
		fmt.Println("go goroutine执行")
		c <- 666 //将666传给c
	}()

	num := <-c
	fmt.Println(num)
	defer fmt.Println("main goroutine 结束")
	fmt.Println("main goroutine执行")

}
