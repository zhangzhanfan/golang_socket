package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	mathRand()
	cryptoRand()
}

// 用math实现，效率快，伪随机数
func mathRand() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(1000)
	fmt.Println("生成一个随机数:", n)
}

// 用
func cryptoRand() {
	by := make([]byte, 8)
	rand.Read(by)
	fmt.Println("生成一个随机数:", by)
}
