/*
	channel的关闭
*/

package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	//方法一
	// for {
	// 	// ok表示
	// 	if data, ok := <-c; ok {
	// 		fmt.Println(data)
	// 	} else {
	// 		break
	// 	}
	// }

	//方法二
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main 关闭")
}
