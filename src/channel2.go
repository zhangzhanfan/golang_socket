/*
	有缓存的channel
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("子go程结束...")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程发送:", i, "len(c)=", len(c), "cap(c)=", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("main接收到 ", num)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("main程结束 ")

}
